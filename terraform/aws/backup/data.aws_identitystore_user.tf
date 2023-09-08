data "aws_identitystore_user" "example" {
  identity_store_id = "d-9367058404" //tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "UserName"
      attribute_value = "ExampleUser"
    }
  }
}

output "user_id" {
  value = data.aws_identitystore_user.example.user_id
}