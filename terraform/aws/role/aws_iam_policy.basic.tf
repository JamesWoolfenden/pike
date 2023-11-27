resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:CreateVpc",
          "ec2:DescribeVpcs",
          "ec2:DescribeVpcAttribute",
          "ec2:DeleteVpc",
          "ec2:CreateSubnet",
          "ec2:DeleteSubnet",
          "ec2:DescribeSubnets",
          "ec2:DescribeNetworkInterfaces",
          "ec2:ModifyVpcAttribute",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface",
          //aws_opensearchserverless_vpc_endpoint
          "aoss:BatchGetVpcEndpoint",
          "aoss:CreateVpcEndpoint",
          "aoss:DeleteVpcEndpoint",
          "aoss:UpdateVpcEndpoint",
          "ec2:CreateTags",
          "ec2:CreateVpcEndpoint",
          "ec2:DeleteNetworkInterface",
          "ec2:DeleteVpcEndpoints",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeVpcEndpoints",
          "ec2:ModifyVpcEndpoints",
          "route53:AssociateVPCWithHostedZone",
          "ec2:DetachNetworkInterface",

          //aws_opensearchserverless_security_config
          "aoss:CreateSecurityConfig",
          "aoss:DeleteSecurityConfig",
          "aoss:GetSecurityConfig",
          "aoss:UpdateSecurityConfig",
          "aoss:DeleteSecurityConfig",

          //aws_opensearchserverless_lifecycle_policy
          "aoss:CreateLifecyclePolicy",
          "aoss:BatchGetLifecyclePolicy",
          "aoss:DeleteLifecyclePolicy",
          "aoss:UpdateLifecyclePolicy",

          //aws_opensearchserverless_collection
          "aoss:CreateCollection",
          "aoss:UpdateCollection",
          "aoss:DeleteCollection",
          "aoss:BatchGetCollection",
          "iam:CreateServiceLinkedRole",
          "aoss:ListTagsForResource",

          //aws_opensearchserverless_access_policy
          "aoss:CreateAccessPolicy",
          "aoss:UpdateAccessPolicy",
          "aoss:DeleteAccessPolicy",

          //aws_opensearchserverless_security_policy
          "aoss:CreateSecurityPolicy",
          "aoss:DeleteSecurityPolicy",
          "aoss:UpdateSecurityPolicy",
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
