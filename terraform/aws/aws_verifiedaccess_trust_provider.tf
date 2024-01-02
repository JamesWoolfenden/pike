resource "aws_verifiedaccess_trust_provider" "pike" {
  policy_reference_name    = "example"
  trust_provider_type      = "user"
  user_trust_provider_type = "iam-identity-center"
}
