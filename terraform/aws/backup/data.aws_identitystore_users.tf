
data "aws_identitystore_users" "pike" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]
}
output "aws_identitystore_users" {
  value = data.aws_identitystore_users.pike
}
