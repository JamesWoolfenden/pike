resource "aws_quicksight_user" "pike" {
  session_name  = "an-author"
  email         = "author@example.com"
  namespace     = aws_quicksight_namespace.pike.namespace
  identity_type = "IAM"
  iam_arn       = "arn:aws:iam::680235478471:user/jameswoolfenden"
  user_role     = "AUTHOR"
}
