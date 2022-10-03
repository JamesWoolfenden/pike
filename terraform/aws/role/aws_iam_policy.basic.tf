resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface",

          "ec2:DescribeAccountAttributes",

          "ec2:DescribeNetworkInsightsAnalyses",
          "ec2:DeleteNetworkInsightsAnalysis",

          "tiros:*",
          "ec2:CreateTags",
          "ec2:DeleteTags",
          "cloudfront:CreateOriginAccessControl",
          "cloudfront:GetOriginAccessControl",
          "ec2:CreateNetworkInsightsAccessScope",
          "ec2:DeleteNetworkInsightsAccessScopeAnalysis",
          "ec2:DeleteNetworkInsightsAccessScope",
          "ec2:DescribeAvailabilityZones",
          "ec2:DescribeCustomerGateways",
          "ec2:DescribeInstances",
          "ec2:DescribeInternetGateways",
          "ec2:DescribeManagedPrefixLists",
          "ec2:DescribeNatGateways",
          "ec2:DescribeNetworkAcls",
          "ec2:DescribeNetworkInsightsAccessScopeAnalyses",
          "ec2:DescribeNetworkInsightsAccessScopes",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribePrefixLists",
          "ec2:DescribeRegions",
          "ec2:DescribeRouteTables",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeTransitGatewayAttachments",
          "ec2:DescribeTransitGatewayConnects",
          "ec2:DescribeTransitGatewayPeeringAttachments",
          "ec2:DescribeTransitGatewayRouteTables",
          "ec2:DescribeTransitGatewayVpcAttachments",
          "ec2:DescribeTransitGateways",
          "cloudfront:DeleteOriginAccessControl",
          "ec2:DescribeVpcPeeringConnections",
          "ec2:DescribeVpcs",
          "ec2:DescribeVpnGateways",
          "ec2:GetNetworkInsightsAccessScopeAnalysisFindings",
          "ec2:GetNetworkInsightsAccessScopeContent",
          "ec2:StartNetworkInsightsAccessScopeAnalysis",
          "elasticloadbalancing:DescribeListeners",
          "elasticloadbalancing:DescribeLoadBalancerAttributes",
          "elasticloadbalancing:DescribeLoadBalancers",
          "elasticloadbalancing:DescribeRules",
          "elasticloadbalancing:DescribeTags",
          "cloudfront:UpdateOriginAccessControl",
          "elasticloadbalancing:DescribeTargetHealth",
          "network-firewall:DescribeFirewall",
          "network-firewall:DescribeFirewallPolicy",
          "network-firewall:ListFirewalls",
          "network-firewall:ListRuleGroups",
          "resource-groups:ListGroupResources",
          "tag:GetResources",
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
