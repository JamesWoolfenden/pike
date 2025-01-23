data "aws_ssoadmin_instances" "example" {}

resource "aws_identitystore_user" "example" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]
  display_name      = "John Doe"
  user_name         = "john.doe@example.com"

  name {
    family_name = "Doe"
    given_name  = "John"
  }
}



resource "aws_identitystore_group_membership" "example" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]
  group_id          = aws_identitystore_group.example.group_id
  member_id         = aws_identitystore_user.example.user_id
}
