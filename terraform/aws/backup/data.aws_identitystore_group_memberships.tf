data "aws_ssoadmin_instances" "example" {}

data "aws_identitystore_group" "example" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "DisplayName"
      attribute_value = "ExampleGroup"
    }
  }
}

data "aws_identitystore_group_memberships" "pike" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]
  group_id          = data.aws_identitystore_group.example.group_id
}
output "aws_identitystore_group_memberships" {
  value = data.aws_identitystore_group_memberships.pike
}
