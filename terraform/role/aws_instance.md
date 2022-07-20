resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:DescribeInstances",
          "ec2:TerminateInstances",
          "ec2:StartInstances",
          "ec2:DescribeTags",
          "ec2:DescribeInstanceAttribute",
          "ec2:DescribeVolumes",
          "ec2:DescribeInstanceTypes",
          "ec2:RunInstances",
          "ec2:ModifyInstanceAttribute",
          "ec2:StopInstances",
        "ec2:DescribeInstanceCreditSpecifications"]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
