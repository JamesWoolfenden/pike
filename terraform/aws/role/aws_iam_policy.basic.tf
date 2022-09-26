resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:CreateVpnGateway",
          #          "ec2:AttachVpnGateway",
          "ec2:DescribeVpnGateways",
          "ec2:DeleteVpnGateway",
          #          "ec2:DetachVpnGateway",
          "ec2:CreateTags",
          "ec2:DeleteTags",
          "ec2:DetachVpnGateway",
          "ec2:DescribeAccountAttributes",
          #          "ec2:CreateRouteTable",
          #          "ec2:DescribeRouteTables",
          #          "ec2:DeleteRouteTable",
          #          "ec2:CreateTags",
          #          "ec2:DeleteTags",
          #          "ec2:DeleteRouteTable",

          #          "ec2:DescribeAccountAttributes",
          #          "ec2:DescribeVpnGateways",
          #          "ec2:EnableVgwRoutePropagation",
          #          "ec2:DisableVgwRoutePropagation",
          #          "ec2:CreateTags",
          #          "ec2:DeleteTags",

          #          "ec2:DescribeAccountAttributes",
          #          "ec2:EnableVgwRoutePropagation",
          #          "ec2:DisableVgwRoutePropagation",
          "ec2:AttachVpnGateway",
          "ec2:DetachVpnGateway",
          "ec2:DescribeVpnGateways",
          "ec2:CreateTags",
          "ec2:DeleteTags",
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
