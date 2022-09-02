resource "aws_iam_group_policy" "example" {
  name  = "my_developer_policy_test"
  group = "test"


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
