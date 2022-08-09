resource "aws_iam_user_policy" "example" {
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:Describe*",
        ]
        Effect   = "Deny"
        Resource = "*"
      },
    ]
  })
  user = "basic"
}
