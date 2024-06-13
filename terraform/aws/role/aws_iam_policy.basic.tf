resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_ec2_client_vpn_authorization_rule
          "ec2:DescribeClientVpnAuthorizationRules",

          //aws_ec2_client_vpn_network_association
          "ec2:AssociateClientVpnTargetNetwork",
          "ec2:DisassociateClientVpnTargetNetwork",

          //aws_ec2_client_vpn_route
          "ec2:CreateClientVpnRoute",
          "ec2:DescribeClientVpnRoutes",
          "ec2:DeleteClientVpnRoute",

          //aws_ec2_host
          "ec2:AllocateHosts",
          "ec2:DescribeHosts",
          "ec2:ReleaseHosts",
          "ec2:CreateTags",
          "ec2:DeleteTags",

          //aws_ec2_client_vpn_endpoint
          "ec2:CreateClientVpnEndpoint",
          "ec2:DeleteClientVpnEndpoint",
          "ec2:DescribeClientVpnEndpoints",
          "ec2:ModifyClientVpnEndpoint",
          "iam:CreateServiceLinkedRole",
          "ec2:CreateTags",
          "ec2:DeleteTags",

          //aws_ec2_carrier_gateway
          "ec2:CreateCarrierGateway",
          "ec2:DeleteCarrierGateway",
          "ec2:DescribeCarrierGateways",
          "ec2:CreateTags",
          "ec2:DeleteTags",

          //aws_ec2_fleet
          "ec2:CreateFleet",
          "ec2:RunInstances",
          "ec2:DescribeFleets",
          "ec2:DeleteFleets",
          "ec2:CreateTags",
          "ec2:DeleteTags",

          "ec2:DescribeNetworkInterfaces",
          "acm:AddTagsToCertificate",
          "acm:DeleteCertificate",
          "acm:DescribeCertificate",
          "acm:ListTagsForCertificate",
          "acm:RemoveTagsFromCertificate",
          "acm:RequestCertificate",
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:CreateLaunchTemplate",
          "ec2:CreateLaunchTemplateVersion",
          "ec2:CreateSubnet",
          "ec2:DeleteLaunchTemplate",
          "ec2:DeleteSubnet",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeInstanceTypes",
          "ec2:DescribeLaunchTemplateVersions",
          "ec2:DescribeLaunchTemplates",
          "ec2:DescribeSubnets",
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:DeleteLogGroup",
          "logs:DeleteLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:ListTagsLogGroup",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject"

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
