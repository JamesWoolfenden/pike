resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          #         "ec2:DescribeDhcpOptions",
          #         "ec2:DescribeVpcs",
          #          "ec2:CreateDefaultVpc",
          #         "ec2:DescribeVpcAttribute",
          #          "ec2:DescribeSecurityGroups",
          #
          "ec2:CreateTags",
          "ec2:DeleteTags",
          "ec2:DescribeSubnets",
          "ec2:DescribeDhcpOptions"


          #          "ec2:DescribeRouteTables",
          #          "ec2:DeleteRoute",
          #
          #          "ec2:DescribeSubnets"

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
