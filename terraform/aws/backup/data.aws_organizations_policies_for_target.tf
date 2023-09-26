data "aws_organizations_policies_for_target" "pike" {
  target_id = data.aws_organizations_organization.pike.roots[0].id
  filter    = "SERVICE_CONTROL_POLICY"
}
