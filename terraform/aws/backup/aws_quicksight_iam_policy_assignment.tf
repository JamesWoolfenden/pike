resource "aws_quicksight_iam_policy_assignment" "pike" {
  assignment_name   = "example"
  assignment_status = "ENABLED"
  policy_arn        = aws_iam_policy.example.arn
  identities {
    user = [aws_quicksight_user.example.user_name]
  }

}
