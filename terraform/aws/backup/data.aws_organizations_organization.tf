data "aws_organizations_organization" "example" {}

output "organization" {
  value = data.aws_organizations_organization.example
}
