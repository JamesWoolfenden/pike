data "aws_organizations_organizational_unit_child_accounts" "pike" {
  parent_id = data.aws_organizations_organization.pike.roots[0].id
}

data "aws_organizations_organization" "pike" {}
