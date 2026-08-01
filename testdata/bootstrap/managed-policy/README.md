# managed-policy fixture

No `terraform/` directory: this fixture uses a synthetic `policy.Policy`
(`iam:CreatePolicyVersion`) defined directly in the test, rather than a
scanned Terraform resource. Pike's `aws_iam_policy` mapping doesn't
naturally produce `CreatePolicyVersion`/`SetDefaultPolicyVersion`/etc.
(those are provider-internal calls on policy update, not something the
static resource mapping models), so constructing the input directly is
more reliable than hunting for a real resource that happens to trigger it.

This fixture demonstrates that deploy-boundary generation denies a
policy-mutation action outright even though the deploy policy "requests"
it — `deploy-policy.json` includes `iam:CreatePolicyVersion`, but
`deploy-boundary.json` has no matching Allow statement anywhere, only the
mandatory `DenyBoundaryPolicyMutation` deny. That's the
"deploy policy = what's requested, deploy boundary = maximum safe
delegation" split (spec section 11) working as intended.
