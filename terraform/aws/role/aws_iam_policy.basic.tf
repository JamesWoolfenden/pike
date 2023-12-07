resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "route53:ListHostedZones",
          "ec2:DescribeVpcs",
          "route53:GetHostedZone",
          "route53:ListTagsForResource",
          "ec2:DescribeVpcAttribute",
          "elasticloadbalancing:CreateLoadBalancer",
          "elasticloadbalancing:AddTags",
          "elasticloadbalancing:RemoveTags",
          "elasticloadbalancing:CreateLoadBalancerListeners",
          "elasticloadbalancing:DescribeLoadBalancers",
          "elasticloadbalancing:DescribeLoadBalancerAttributes",
          "ec2:DescribeSecurityGroups",
          "elasticloadbalancing:DescribeTags",
          "elasticloadbalancing:DeleteLoadBalancer",
          "elasticloadbalancing:ModifyLoadBalancerAttributes",
          "ec2:DescribeNetworkInterfaces",
          "elasticloadbalancing:EnableAvailabilityZonesForLoadBalancer",
          "ec2:DetachNetworkInterface",

          //aws_route53_traffic_policy
          "route53:CreateTrafficPolicy",
          "route53:GetTrafficPolicy",
          "route53:DeleteTrafficPolicy",
          "route53:ListTrafficPolicies",
          "route53:ListTrafficPolicyVersions",

          //aws_route53_resolver_rule
          "route53resolver:CreateResolverRule",
          "route53resolver:GetResolverRule",
          "route53resolver:DeleteResolverRule",
          "route53resolver:UpdateResolverRule",
          "route53resolver:TagResource",
          "route53resolver:UntagResource",

          //aws_route53_resolver_firewall_rule_group
          "route53resolver:CreateFirewallRuleGroup",
          "route53resolver:GetFirewallRuleGroup",
          "route53resolver:DeleteFirewallRuleGroup",
          "route53resolver:ListTagsForResource",
          "route53resolver:TagResource",
          "route53resolver:UntagResource",

          //aws_route53_resolver_firewall_domain_list
          "route53resolver:CreateFirewallDomainList",
          "route53resolver:GetFirewallDomainList",
          "route53resolver:DeleteFirewallDomainList",
          "route53resolver:ListFirewallDomains",
          "route53resolver:TagResource",
          "route53resolver:UntagResource",

          //aws_route53_resolver_firewall_config
          "route53resolver:UpdateFirewallConfig",
          "route53resolver:GetFirewallConfig",
          "route53resolver:ListFirewallConfigs",

          //aws_route53_resolver_dnssec_config
          "route53resolver:UpdateResolverDnssecConfig",
          "route53resolver:GetResolverDnssecConfig",
          "route53resolver:ListResolverDnssecConfigs",

          //aws_route53_resolver_config
          "route53resolver:UpdateResolverConfig",
          "route53resolver:GetResolverConfig",
          "route53resolver:ListResolverConfigs",

          //aws_route53_cidr_collection
          "route53:CreateCidrCollection",
          "route53:ChangeCidrCollection",
          "route53:DeleteCidrCollection",
          "route53:ListCidrCollections",

          //aws_route53_resolver_firewall_rule_group_association
          "route53resolver:AssociateFirewallRuleGroup",
          "route53resolver:GetFirewallRuleGroupAssociation",
          "route53resolver:DisassociateFirewallRuleGroup",
          "route53resolver:UpdateFirewallRuleGroupAssociation",

          //aws_route53_resolver_firewall_rule
          "route53resolver:CreateFirewallRule",
          "route53resolver:ListFirewallRules",
          "route53resolver:DeleteFirewallRule",
          "route53resolver:UpdateFirewallRule",

          //aws_route53_cidr_location
          "route53:ListCidrBlocks",

          //aws_route53_traffic_policy_instance
          "route53:CreateTrafficPolicyInstance",
          "route53:GetTrafficPolicyInstance",
          "route53:DeleteTrafficPolicyInstance",

          //aws_load_balancer_policy
          "elasticloadbalancing:CreateLoadBalancerPolicy",
          "elasticloadbalancing:DescribeLoadBalancerPolicies",
          "elasticloadbalancing:DeleteLoadBalancerPolicy",
          "elasticloadbalancing:AddTags",
          "elasticloadbalancing:RemoveTags",

          //aws_load_balancer_listener_policy
          "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",

          //aws_load_balancer_backend_server_policy
          "elasticloadbalancing:SetLoadBalancerPoliciesForBackendServer",

          //aws_route53_resolver_endpoint

          "route53resolver:ListResolverEndpoints",
          "route53resolver:GetResolverEndpoint",
          "route53resolver:DeleteResolverEndpoint",
          "route53resolver:UpdateResolverEndpoint",
          "route53resolver:CreateResolverEndpoint",
          "route53resolver:TagResource",
          "route53resolver:UntagResource",

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
