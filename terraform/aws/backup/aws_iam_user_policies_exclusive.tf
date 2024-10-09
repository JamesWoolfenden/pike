resource "aws_iam_user_policies_exclusive" "pike" {
  user_name    = aws_iam_user.pike.name
  policy_names = [aws_iam_policy.demo.name]
}


resource "aws_iam_user" "pike" {
  name = "pike"
}
