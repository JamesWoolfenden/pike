resource "aws_iam_group_policy_attachment" "example" {
  group      = "test"
  policy_arn = "arn:aws:iam::680235478471:policy/basic"
}
