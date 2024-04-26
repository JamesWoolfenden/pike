data "aws_identitystore_groups" "pike" {
  identity_store_id = data.aws_ssoadmin_instances.example.identity_store_ids[0]

}

output "aws_identitystore_groups" {
  value = data.aws_identitystore_groups.pike
}

data "aws_ssoadmin_instances" "example" {}
