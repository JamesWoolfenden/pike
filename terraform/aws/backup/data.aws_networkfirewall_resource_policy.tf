data "aws_networkfirewall_resource_policy" "pike" {
  resource_arn = aws_networkfirewall_resource_policy.example.resource_arn
}
#
#resource "aws_networkfirewall_resource_policy" "example" {
#  resource_arn = aws_networkfirewall_firewall_policy.example.arn
#  # policy's Action element must include all of the following operations
#  policy = jsonencode({
#    Statement = [{
#      Action = [
#        "network-firewall:ListFirewallPolicies",
#        "network-firewall:CreateFirewall",
#        "network-firewall:UpdateFirewall",
#        "network-firewall:AssociateFirewallPolicy"
#      ]
#      Effect   = "Allow"
#      Resource = aws_networkfirewall_firewall_policy.example.arn
#      Principal = {
#        AWS = "arn:aws:iam::123456789012:root"
#      }
#    }]
#    Version = "2012-10-17"
#  })
#}
#
#resource "aws_networkfirewall_firewall_policy" "example" {
#  name = "example"
#
#  firewall_policy {
#    stateless_default_actions          = ["aws:pass"]
#    stateless_fragment_default_actions = ["aws:drop"]
#    stateless_rule_group_reference {
#      priority     = 1
#      resource_arn = aws_networkfirewall_rule_group.example.arn
#    }
#  }
#
#  tags = {
#    Tag1 = "Value1"
#    Tag2 = "Value2"
#  }
#}
#
#
#resource "aws_networkfirewall_rule_group" "example" {
#  capacity = 100
#  name     = "example"
#  type     = "STATEFUL"
#  rule_group {
#    rules_source {
#      rules_source_list {
#        generated_rules_type = "DENYLIST"
#        target_types         = ["HTTP_HOST"]
#        targets              = ["test.example.com"]
#      }
#    }
#  }
#
#  tags = {
#    Tag1 = "Value1"
#    Tag2 = "Value2"
#  }
#}
