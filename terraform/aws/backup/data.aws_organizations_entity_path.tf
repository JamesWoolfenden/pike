data "aws_organizations_entity_path" "pike" {
  entity_id = "ou-ghi0-awsccccc"
}

output "aws_organizations_entity_path" {
  value = data.aws_organizations_entity_path.pike
}
