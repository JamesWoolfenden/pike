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

Every Resource and Condition you add must be traceable to something literal in the Terraform, or to a real datasource reference Terraform will resolve at apply time (see step 2). Where a name is genuinely unresolvable either way (computed at apply time via `random_id`/`random_pet`, or supplied only through a variable with no default), leave that specific ARN segment as `*` and say so explicitly in the rationale — a guessed ARN that's wrong is worse than an honest wildcard, because it fails closed silently instead of visibly.

---

## Process

1. **Get the deterministic floor.** Run:
   - `! pike scan -o split -d <dir>` — base/escalation action split per provider (`SplitPolicy` JSON: `{"aws":{"base":[...],"escalation":[...]}, ...}`)
   - `! pike scan -o json -d <dir>` — the full `Sorted` bag, including `PlanAWS`/`PlanGCP`/`PlanAZURE` (planner-role actions) and `RuntimeAWS`/`RuntimeGCP`/`RuntimeAZURE` (per-resource runtime permissions, already resource-attributed)
   These three lists are ground truth. Every action in your two output policies must come from here — nothing added, nothing removed.

2. **Read the Terraform.** For every resource that contributed actions to the bag above, find its declaration and pull out anything ARN-able: literal bucket/table/role/key/queue names, `name_prefix`, explicit `arn` references to sibling resources declared in the same config, account/region if hardcoded via `provider` blocks. Build a resource → ARN-pattern map. Where a resource name is computed (e.g. `"${var.env}-${random_id.suffix.hex}"`), use the literal parts and wildcard only the unresolvable segment — don't wildcard the whole ARN just because part of it is dynamic.

   **Prefer a datasource reference over a wildcard for account/project/subscription segments.** The output of this skill is Terraform, not static JSON — so an account ID, project ID, or subscription ID that can't be known statically doesn't have to become `*`. It can be a real interpolation that Terraform resolves correctly at apply time: `data.aws_caller_identity.current.account_id` (AWS), `data.google_project.current.project_id` (GCP), `data.azurerm_client_config.current.subscription_id` (Azure). Check whether the scanned Terraform already declares the equivalent datasource block and reuse that reference rather than declaring a duplicate; otherwise add the datasource block to the generated policy's own `.tf` output and note that it's new. Only fall back to a literal `*` for a segment that no well-known "current context" datasource can resolve either (e.g. a cross-account target ARN supplied only via a variable with no default, or a name that's genuinely random at apply time via `random_id`/`random_pet`).

3. **Scope each statement's Resource.** Group actions by the resource(s) they were attributed to (pike's `RuntimePermission.ResourceType`/`ResourceName` already gives you this attribution for runtime permissions; for apply/plan actions, re-derive it from which Terraform block required which action during step 2). Emit one statement per resource-group with `Resource` set to the ARN pattern(s) from step 2, not `["*"]`. Actions that are inherently account/region-scoped only (no per-resource ARN in that service, e.g. `iam:ListRoles`) legitimately keep `Resource: ["*"]` — say so, don't leave it unexplained.

4. **Add Conditions to escalation-class actions.** For every action in the `escalation` list from step 1, determine *why* it's needed from the Terraform (e.g. `iam:PassRole` because a resource passes a role to a specific service principal in its `assume_role_policy`) and attach the narrowest Condition that still lets the real Terraform apply succeed — using the same key families `audit_aws.go`'s `escalationScopingConditionKeys` already treats as genuinely scoping (`iam:PassedToService`, `iam:ResourceTag`, `aws:SourceArn`, `aws:PrincipalTag`, etc. — check that list for the current, authoritative set; it evolves), not an arbitrary condition that merely exists. Mirror the equivalent for GCP (`iam.serviceAccounts.*` scoped to a specific SA resource, not project-wide) and Azure (role assignment `scope` narrowed to the specific resource ID, not the subscription). If no real scoping condition is derivable from the Terraform, say so and leave the action unconditioned rather than inventing a condition that doesn't actually restrict anything — a fake Condition is worse than none, because it reads as safe on review.

5. **Assemble and validate.** Write each candidate policy as a temporary `aws_iam_policy`/`google_project_iam_custom_role`/`azurerm_role_definition` HCL fixture (matching the shape pike's audit handlers already parse — see `auditAWSInlinePolicy`, `auditAzureRoleDefinition`, GCP equivalents), including any datasource blocks from step 2's interpolated Resources, into a scratch directory, then run `! pike audit -d <scratch-dir> -o json -min-severity low`. Any finding means the policy you generated is itself unsafe by pike's own deterministic rules — that's a bug in your scoping, not a false positive to argue past. Tighten and re-run. Cap at 3 iterations; if still unclean, report exactly which finding won't resolve and why (e.g. "AWS004 still fires on `sts:AssumeRole` — no Terraform-derivable principal to scope `Resource` to, this account assumes into an external ARN supplied at runtime").

   Once validation passes, the scratch copy is a draft, not the deliverable — a file in a session-scoped temp directory is invisible to the user the moment the session ends. Copy the final, passing files into the scanned repo (default: `<scanned-dir>/pike-policy/planner.tf` and `applier.tf`, alongside a `variables.tf` for the principal inputs) unless the user names a different location. Only leave them in scratch if the user explicitly says not to write into the repo.

   **Know the checker's blind spot.** `audit_aws.go`'s AWS003 (Resource `"*"` with resource-scopable actions) only fires on a *literal* `"*"` — a Resource string containing `${data.aws_caller_identity.current.account_id}` interpolation gets silently dropped by pike's static HCL parser and reads as "not wildcard, nothing to flag." That's a parser limitation, not confirmation your scoping is correct. AWS004 (the Condition check) is unaffected — it only inspects Action and Condition, never Resource content, so it still fires correctly regardless of interpolation. So: a clean `pike audit` run proves your Conditions are real; it does **not** prove an interpolated Resource is well-formed. State in the report which Resources rely on datasource interpolation and were therefore not independently re-verified by the audit step, as distinct from Resources pike's parser actually checked.

6. **Sanity-check against the floor.** Confirm every action in your two generated policies appears in step 1's lists, and that the planner policy's action set is a subset of the applier's. If either check fails, you introduced or dropped an action while scoping — fix the generation, not the check.

7. **Look for a repeatable pattern.** If the same scoping logic (e.g. "every `iam:PassRole` in this codebase should condition on `iam:PassedToService`") would apply to *any* Terraform using this resource shape, not just this run — that's a candidate for a real deterministic audit rule, not something to re-derive by hand every invocation. Pike's audit rules live in this same repo (`src/audit_aws.go`/`audit_gcp.go`/`audit_azure.go`), so — unlike holden's cross-org fork — there's no PR-authorization barrier here. Propose the specific rule (which file, what check, what severity) and, if the user agrees, implement it with a testdata fixture and a `TestAudit_*` case following the existing pattern, the same way the AWS004/AZURE002 fixes were done.

8. **Report.** The most important thing the caller needs is where the output lives — never make them ask. State the full absolute path to every file written (planner/applier/variables), and show the policy content inline in the response too (both, not one or the other — a path with nothing to preview is as unhelpful as a wall of HCL with no way to find it again later). Alongside that: for each provider, a table of every Resource scoped down from `*` (old → new, with the Terraform reference that justified it), every Condition added (action → condition → why), every action that legitimately keeps `Resource: "*"` or stays unconditioned (with reason), and the final `pike audit` result on the applier policy. If step 7 produced a suggested rule, note it separately as a follow-up, not blocking this run's output.

---

## Discipline

- **Never remove or substitute an action pike listed.** If an action looks wrong, that's a pike bug to report separately — this skill only adds restriction, never edits the action set.
- **Never fabricate an ARN or account ID.** Prefer a datasource reference (`data.aws_caller_identity.current.account_id` and equivalents — see step 2) over a literal guess. Only fall back to wildcarding the segment when no datasource can resolve it either, and say why.
- **Never add a Condition that doesn't actually restrict anything** just to make `pike audit` stop complaining — that defeats the entire point and is worse than the unconditioned finding, because it looks fixed on review.
- **The validation loop (step 5) is mandatory, not optional.** A hand-scoped policy that hasn't been re-run through `pike audit` is a draft, not an output. But know what it did and didn't check: it verifies Conditions on escalation actions reliably; it does **not** verify a datasource-interpolated Resource is well-formed (see step 5's blind-spot note) — don't report "audit passed" as if it covered both.
- **When genuinely stuck** (no derivable scoping, an iteration cap hit, a resource type pike/audit doesn't know), say so plainly in the report rather than emitting a plausible-looking but unverified policy.

---

## Status

Validated on a single hand-crafted AWS fixture (S3 bucket + IAM role assumed by a Lambda + the Lambda function) — the deterministic-floor extraction, ARN scoping, Condition derivation, and the `pike audit` validation loop all worked, including correctly refusing to fabricate a Condition for `iam:PutRolePolicy` when none was derivable and reporting that honestly instead. A follow-up test confirmed the checker's blind spot documented in step 5: swapping a literal Resource for a `${data.aws_caller_identity.current.account_id}` interpolation makes AWS003 (wildcard check) go silent either way, while AWS004 (Condition check) still fires correctly — so a clean audit run only certifies Conditions, not interpolated Resources.

**Not yet exercised**: GCP/Azure targets, computed/interpolated resource names beyond the one datasource case above, multi-resource-of-same-type attribution, and the step 7 "propose a new audit rule" path. Treat output as a strong draft to review, not a policy to attach unreviewed.
