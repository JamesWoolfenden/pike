resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:CreateIpam",
          "ec2:CreateIpamPool",
          "ec2:CreateSecurityGroup",
          "ec2:CreateVPC",
          "ec2:CreateVpcEndpoint",
          "ec2:DeleteIpam",
          "ec2:DeleteIpamPool",
          "ec2:DeleteSecurityGroup",
          "ec2:DeleteVPC",
          "ec2:DeleteVpcEndpoints",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeIpamPools",
          "ec2:DescribeIpams",
          "ec2:DescribeNetworkAcls",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribePrefixLists",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeTags",
          "ec2:DescribeVpcAttribute",
          "ec2:DescribeVpcEndpointServices",
          "ec2:DescribeVpcEndpoints",
          "ec2:DescribeVpcs",
          "ec2:ModifyIpam",
          "ec2:ModifyIpamPool",
          "ec2:ModifyVpcAttribute",
          "ec2:ModifyVpcEndpoint",
          "ec2:ModifyVpcTenancy",
          "ec2:RevokeSecurityGroupEgress",
          # "iam:CreateServiceLinkedRole",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",
          "vpc-lattice:CreateServiceNetworkVpcEndpointAssociation",
          "vpc-lattice:DescribeServiceNetworkVpcEndpointAssociation",
          "ec2:DescribeIpamScopes",
          "ec2:AssociateSecurityGroupVpc",
          "ec2:AssociateVpcCidrBlock",
          "ec2:AllocateIpamPoolCidr",
          "ec2:DescribeSecurityGroupVpcAssociations",
          "ec2:DisassociateSecurityGroupVpc",

          "ec2:ModifyVpcEndpointServicePermissions"

        ],
        "Resource" : [
          "*"
        ]
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
