# unknown-iam-action fixture

No `terraform/` directory: this fixture uses a synthetic `policy.Policy`
(`iam:UpdateAssumeRolePolicy`, an IAM write the classifier doesn't
recognize) defined directly in the test, rather than a scanned Terraform
resource — no readily available Pike-mapped resource produces exactly
this action on its own, and constructing the input directly keeps the
fixture precisely targeted at the fail-closed path.

Generation is expected to fail at the deploy-boundary stage
(`*transform.UnknownIAMActionError`), so `expected/` only contains
`deploy-policy.json` — the one artifact that gets rendered before the
pipeline stops. There is no `deploy-boundary.json`, `workload-boundary.json`,
`trust-policy.json`, `assume-role-policy.json` or `report.json` for this
fixture.
