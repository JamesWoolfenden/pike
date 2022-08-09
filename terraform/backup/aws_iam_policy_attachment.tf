resource "aws_iam_policy_attachment" "example" {
  name       = "test-attachment"
  users      = ["basic"]
  roles      = ["test"]
  groups     = ["test"]
  policy_arn = "arn:aws:iam::680235478471:policy/test"
}
