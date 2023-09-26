data "aws_organizations_policies" "pike" {
  filter = "SERVICE_CONTROL_POLICY"
}
