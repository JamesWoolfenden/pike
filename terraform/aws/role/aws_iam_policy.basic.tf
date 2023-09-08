resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          #          //aws_networkfirewall_firewall_policy
          #          "network-firewall:DescribeFirewallPolicy",
          #          //aws_networkfirewall_firewall
          #          "network-firewall:DescribeFirewall",
          #          //aws_licensemanager_received_licenses, aws_licensemanager_received_license
          #          "license-manager:ListReceivedLicenses",
          #          //aws_licensemanager_grants
          #          "license-manager:ListDistributedGrants"

          //aws_networkfirewall_rule_group
          "network-firewall:CreateRuleGroup",
          "network-firewall:DescribeRuleGroup",
          "network-firewall:DeleteRuleGroup",
          "network-firewall:TagResource",
          "network-firewall:UntagResource",
          "network-firewall:UpdateRuleGroup",

          //aws_networkfirewall_firewall_policy
          "network-firewall:CreateFirewallPolicy",
          "network-firewall:TagResource",
          "network-firewall:UntagResource",
          "network-firewall:ListRuleGroups",
          "network-firewall:DescribeFirewallPolicy",
          "network-firewall:DeleteFirewallPolicy",
          "network-firewall:UpdateFirewallPolicy",

          //aws_networkfirewall_resource_policy
          "network-firewall:PutResourcePolicy",
          "network-firewall:DescribeResourcePolicy",
          "network-firewall:DeleteResourcePolicy"
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
