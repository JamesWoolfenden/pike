data "aws_organizations_organization" "org" {}

data "aws_organizations_organizational_unit_descendant_organizational_units" "pike" {
  parent_id = data.aws_organizations_organization.org.roots[0].id
}

output "aws_organizations_organizational_unit_descendant_organizational_units" {
  value = data.aws_organizations_organizational_unit_descendant_organizational_units.pike
}
