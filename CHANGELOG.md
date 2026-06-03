# Changelog

## v1.0.0 (2026-06-03)

### Breaking changes

- **`pike scan` default output is now two-role** — both an apply role (full permissions for `terraform apply`) and a plan role (read-only permissions for `terraform plan`) are emitted by default. Use `--legacy` to restore the previous single-role output.
- **`pike compare` signature updated** — `strict bool` parameter added. Existing integrations calling the Go library directly must add `false` as the final argument.

### New features

#### Two-role output (apply + plan)

`pike scan` now outputs two permission sets by default, implementing the two-principal split pattern:

- **Apply role** — all permissions needed for `terraform apply`. Bind this to the applier SA on protected branches only.
- **Plan role** — read-only permissions sufficient for `terraform plan`. Bind this to the planner SA on every branch.

Plan permissions are derived from the `plan` arrays in AWS mapping files, and by filtering GCP/Azure apply permissions to read-side verbs (`*.get`, `*.list`, `*/read`).

Use `--legacy` to get the old single-role output.

#### Write mode two-role Terraform (`--write`)

`pike scan --write` now produces `pike.two-role.tf` — a single Terraform file containing both the apply and plan roles with their policies and attachments ready to deploy. In `--legacy` mode, the previous separate files are written instead.

#### `--output split`

`pike scan --output split` partitions permissions into `base` (safe to distribute broadly) and `escalation` (owner-equivalent) subsets per provider:

```json
{
    "aws": {
        "base": ["s3:GetObject", "s3:ListBucket", ...],
        "escalation": ["iam:PassRole", "iam:PutRolePolicy"]
    }
}
```

#### Escalation-class permission warnings

Pike always warns on stderr when the computed role contains owner-equivalent permissions — permissions that allow the holder to grant themselves additional access. No flag is needed; the warning fires automatically.

```text
WARNING: escalation-class permissions detected — holder can grant themselves additional access.
         Consider the two-role pattern: planner SA (read-only) on all branches,
         applier SA (full) on protected branches only.
  Escalation permissions:
    aws:   iam:PassRole, iam:PutRolePolicy
    gcp:   resourcemanager.projects.setIamPolicy
```

#### `pike compare --strict`

Exits non-zero when escalation-class permissions are present in the computed role, even if the deployed and computed policies otherwise match. Useful for enforcing the two-role pattern in CI.

```bash
pike compare --strict -d ./terraform -a arn:aws:iam::123456789012:policy/terraform_pike
```

### Improvements

- `--init` failures now emit a single warning line instead of the full terraform error dump, keeping output readable when scanning directories with conflicting example files.

### Escalation-class permission lists

The built-in escalation sets cover:

**GCP**: `resourcemanager.{projects,folders,organizations}.setIamPolicy`, `iam.roles.{create,update,delete}`, `iam.serviceAccounts.{setIamPolicy,getAccessToken,signBlob,signJwt,implicitDelegation}`, `iam.serviceAccountKeys.create`

**AWS**: `iam:{PutRolePolicy,AttachRolePolicy,CreatePolicyVersion,SetDefaultPolicyVersion,UpdateAssumeRolePolicy,PassRole,CreateAccessKey}`, `sts:AssumeRole`

**Azure**: `Microsoft.Authorization/roleAssignments/write`, `Microsoft.Authorization/roleDefinitions/write`, `Microsoft.ManagedIdentity/userAssignedIdentities/assign/action`
