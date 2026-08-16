---
name: pike-safe-policy
description: pike++ — AI least-privilege layer on top of pike's deterministic permission scan. Adds Resource ARN scoping and Conditions to escalation-class actions, splits base/escalation into the two-role pattern, and validates the result against pike's own audit rules until clean.
---

# pike++

`pike scan` is the deterministic engine: given Terraform, it produces the exact, correct set of IAM **actions** required — no more, no less. That part is a solved problem and this skill must never second-guess it. What pike cannot do is decide **who those actions should apply to** and **under what conditions** — that requires reading the actual resource names in the config and reasoning about intent, which is exactly what a rules engine structurally can't do. `pike++` is you: the layer that turns pike's action list into an actual least-privilege policy.

**Arguments:**

- Optional: a directory to scan (defaults to `.`)
- Usage: `/pike-safe-policy` or `/pike-safe-policy ./terraform`

---

## The gap, precisely

Don't take this on faith — it's visible in pike's own source:

- `src/types.go` — `Statement` has `Sid`, `Effect`, `Action`, `Resource` only. **No `Condition` field exists in the type.** pike cannot emit a conditioned statement even in principle, today.
- `src/policy.go:58-60` — `myResource := []string{"*"}` / `resource := "*"` is hardcoded before the ARN builder ever runs. Every generated statement's Resource is `"*"`, regardless of what the Terraform actually names.
- `src/escalation.go` (`escalationAWS`/`escalationGCP`/`escalationAZURE`) and `WarnEscalation` already identify *which* actions are escalation-class and recommend a two-role split — but pike only warns; it doesn't produce the two policy documents.
- `src/audit_aws.go` / `audit_gcp.go` / `audit_azure.go` already know what a *safe* Condition looks like (see `escalationScopingConditionKeys` in `audit_aws.go`) — that judgment exists in the codebase for reviewing *hand-written* policies, but nothing wires it into *generation*.

So: pike tells you the floor (never remove an action pike says is required). This skill's only job is to add restriction on top — Resource scoping and Conditions — never to narrow the action list itself. If you find yourself dropping an action pike listed because "it looks too broad," stop — that's a bug in this skill, not a fix.

---

## What you produce

Two IAM policy documents per provider, mirroring pike's own two-role recommendation:

1. **Planner** — `PlanAWS`/`PlanGCP`/`PlanAZURE` actions only (read-only, safe for `terraform plan` on every branch). Never contains an escalation-class action.
2. **Applier** — the full `Sorted.AWS`/`GCP`/`AZURE` action list (everything, including escalation-class), scoped down with Resource ARNs and Conditions. Restricted to protected branches / a gated role.

Every Resource and Condition you add must be traceable to something literal in the Terraform. Where a name is genuinely unresolvable statically (computed, interpolated from a data source, account ID unknown until apply), leave that specific ARN segment as `*` and say so explicitly in the rationale — a guessed ARN that's wrong is worse than an honest wildcard, because it fails closed silently instead of visibly.

---

## Process

1. **Get the deterministic floor.** Run:
   - `! pike scan -o split -d <dir>` — base/escalation action split per provider (`SplitPolicy` JSON: `{"aws":{"base":[...],"escalation":[...]}, ...}`)
   - `! pike scan -o json -d <dir>` — the full `Sorted` bag, including `PlanAWS`/`PlanGCP`/`PlanAZURE` (planner-role actions) and `RuntimeAWS`/`RuntimeGCP`/`RuntimeAZURE` (per-resource runtime permissions, already resource-attributed)
   These three lists are ground truth. Every action in your two output policies must come from here — nothing added, nothing removed.

2. **Read the Terraform.** For every resource that contributed actions to the bag above, find its declaration and pull out anything ARN-able: literal bucket/table/role/key/queue names, `name_prefix`, explicit `arn` references to sibling resources declared in the same config, account/region if hardcoded via `provider` blocks or `data.aws_caller_identity`/`data.aws_region` references. Build a resource → ARN-pattern map. Where a resource name is computed (e.g. `"${var.env}-${random_id.suffix.hex}"`), use the literal parts and wildcard only the unresolvable segment — don't wildcard the whole ARN just because part of it is dynamic.

3. **Scope each statement's Resource.** Group actions by the resource(s) they were attributed to (pike's `RuntimePermission.ResourceType`/`ResourceName` already gives you this attribution for runtime permissions; for apply/plan actions, re-derive it from which Terraform block required which action during step 2). Emit one statement per resource-group with `Resource` set to the ARN pattern(s) from step 2, not `["*"]`. Actions that are inherently account/region-scoped only (no per-resource ARN in that service, e.g. `iam:ListRoles`) legitimately keep `Resource: ["*"]` — say so, don't leave it unexplained.

4. **Add Conditions to escalation-class actions.** For every action in the `escalation` list from step 1, determine *why* it's needed from the Terraform (e.g. `iam:PassRole` because a resource passes a role to a specific service principal in its `assume_role_policy`) and attach the narrowest Condition that still lets the real Terraform apply succeed — using the same key families `audit_aws.go`'s `escalationScopingConditionKeys` already treats as genuinely scoping (`iam:PassedToService`, `iam:ResourceTag`, `aws:SourceArn`, `aws:PrincipalTag`, etc. — check that list for the current, authoritative set; it evolves), not an arbitrary condition that merely exists. Mirror the equivalent for GCP (`iam.serviceAccounts.*` scoped to a specific SA resource, not project-wide) and Azure (role assignment `scope` narrowed to the specific resource ID, not the subscription). If no real scoping condition is derivable from the Terraform, say so and leave the action unconditioned rather than inventing a condition that doesn't actually restrict anything — a fake Condition is worse than none, because it reads as safe on review.

5. **Assemble and validate.** Write each candidate policy as a temporary `aws_iam_policy`/`google_project_iam_custom_role`/`azurerm_role_definition` HCL fixture (matching the shape pike's audit handlers already parse — see `auditAWSInlinePolicy`, `auditAzureRoleDefinition`, GCP equivalents) into a scratch directory, then run `! pike audit -d <scratch-dir> -o json -min-severity low`. Any finding means the policy you generated is itself unsafe by pike's own deterministic rules — that's a bug in your scoping, not a false positive to argue past. Tighten and re-run. Cap at 3 iterations; if still unclean, report exactly which finding won't resolve and why (e.g. "AWS004 still fires on `sts:AssumeRole` — no Terraform-derivable principal to scope `Resource` to, this account assumes into an external ARN supplied at runtime").

6. **Sanity-check against the floor.** Confirm every action in your two generated policies appears in step 1's lists, and that the planner policy's action set is a subset of the applier's. If either check fails, you introduced or dropped an action while scoping — fix the generation, not the check.

7. **Look for a repeatable pattern.** If the same scoping logic (e.g. "every `iam:PassRole` in this codebase should condition on `iam:PassedToService`") would apply to *any* Terraform using this resource shape, not just this run — that's a candidate for a real deterministic audit rule, not something to re-derive by hand every invocation. Pike's audit rules live in this same repo (`src/audit_aws.go`/`audit_gcp.go`/`audit_azure.go`), so — unlike holden's cross-org fork — there's no PR-authorization barrier here. Propose the specific rule (which file, what check, what severity) and, if the user agrees, implement it with a testdata fixture and a `TestAudit_*` case following the existing pattern, the same way the AWS004/AZURE002 fixes were done.

8. **Report.** For each provider: the planner and applier policy JSON, a table of every Resource scoped down from `*` (old → new, with the Terraform reference that justified it), every Condition added (action → condition → why), every action that legitimately keeps `Resource: "*"` or stays unconditioned (with reason), and the final `pike audit` result on the applier policy. If step 7 produced a suggested rule, note it separately as a follow-up, not blocking this run's output.

---

## Discipline

- **Never remove or substitute an action pike listed.** If an action looks wrong, that's a pike bug to report separately — this skill only adds restriction, never edits the action set.
- **Never fabricate an ARN or account ID.** Wildcard the part you can't resolve statically and say why, rather than guessing a plausible-looking value.
- **Never add a Condition that doesn't actually restrict anything** just to make `pike audit` stop complaining — that defeats the entire point and is worse than the unconditioned finding, because it looks fixed on review.
- **The validation loop (step 5) is mandatory, not optional.** A hand-scoped policy that hasn't been re-run through `pike audit` is a draft, not an output.
- **When genuinely stuck** (no derivable scoping, an iteration cap hit, a resource type pike/audit doesn't know), say so plainly in the report rather than emitting a plausible-looking but unverified policy.
