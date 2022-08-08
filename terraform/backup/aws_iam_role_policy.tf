resource "aws_iam_role_policy" "example" {

  name = "test_policy"
  role = "pike-test-role"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:Describe*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })

}
