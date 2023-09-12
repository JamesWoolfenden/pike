#data "aws_ssoadmin_instances" "example" {}

data "aws_identitystore_group" "example" {
  identity_store_id = "d-9367058404" //tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "DisplayName"
      attribute_value = "ExampleGroup"
    }
  }
}

output "group_id" {
  value = data.aws_identitystore_group.example.group_id
}
