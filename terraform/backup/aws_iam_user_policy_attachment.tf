resource "aws_iam_user_policy_attachment" "test-attach" {
  user       = "basic"
  policy_arn = "arn:aws:iam::680235478471:policy/Terraform"
}
