resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:DescribeAccountAttributes",
          "lambda:CreateEventSourceMapping",
          "lambda:GetEventSourceMapping",
          "lambda:DeleteEventSourceMapping",
          "lambda:UpdateEventSourceMapping",

          "ec2:DescribeAccountAttributes",
          "lambda:UpdateFunctionEventInvokeConfig",
          "lambda:GetFunctionEventInvokeConfig",
          "lambda:PutFunctionEventInvokeConfig",
          "lambda:DeleteFunctionEventInvokeConfig",

          "ec2:DescribeAccountAttributes",
          "lambda:CreateFunctionUrlConfig",
          "lambda:GetFunctionUrlConfig",
          "lambda:DeleteFunctionUrlConfig",
          "lambda:UpdateFunctionUrlConfig",
          "lambda:AddPermission",

          "ec2:DescribeAccountAttributes",
          "lambda:PutProvisionedConcurrencyConfig",
          "lambda:DeleteProvisionedConcurrencyConfig",
          "lambda:GetProvisionedConcurrencyConfig",

        ]
        "Resource" : "*"
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
