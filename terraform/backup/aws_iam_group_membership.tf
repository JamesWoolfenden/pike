resource "aws_iam_group_membership" "example" {
  name = "muddle"
  users = [
    "karthik",
    //"jameswoolfenden"
  ]

  group = "test"
}
