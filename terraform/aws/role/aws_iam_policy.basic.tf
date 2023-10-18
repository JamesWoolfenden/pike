resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //          "sqs:*",
          "cloudformation:DescribeStackEvents",
          "cloudformation:DescribeStackResources",
          "cloudformation:ValidateTemplate",
          //aws_cloudformation_stack
          "cloudformation:CreateStack",
          "cloudformation:DescribeStacks",
          "cloudformation:GetTemplate",
          "cloudformation:DeleteStack",
          "cloudformation:UpdateStack",
          //aws_cloudformation_type
          "iam:PassRole",
          "cloudformation:RegisterType",
          "cloudformation:DescribeType",

          //
          "ec2:CreateVpc",
          "ec2:CreateTags",
          "ec2:DeleteTags",
          "ec2:DescribeVpcs",
          "ec2:DeleteVpc",
          "ec2:ModifyVpcAttribute"

        ],
        "Resource" : "*",
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
