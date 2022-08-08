resource "aws_iam_role_policy_attachment" "example" {
  role       = "lambda_basic"
  policy_arn = "arn:aws:iam::680235478471:policy/basic"
}
