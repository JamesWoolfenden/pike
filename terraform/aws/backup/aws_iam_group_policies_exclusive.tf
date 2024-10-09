resource "aws_iam_group_policies_exclusive" "pike" {
  group_name   = "pike"
  policy_names = [aws_iam_policy.demo.name]
}


resource "aws_iam_policy" "demo" {
  name        = "test_policy"
  path        = "/"
  description = "My test policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
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
