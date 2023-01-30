resource "aws_iam_user_group_membership" "pike" {
  user = "basic"

  groups = [
    "test",
    "pike"
  ]
}