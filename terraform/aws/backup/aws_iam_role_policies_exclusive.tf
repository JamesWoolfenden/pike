resource "aws_iam_role_policies_exclusive" "pike" {
  policy_names = [aws_iam_policy.demo.name]
  role_name    = aws_iam_role.pike.name
}

resource "aws_iam_role" "pike" {
  name = "pike"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}
