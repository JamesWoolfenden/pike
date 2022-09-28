resource "aws_networkfirewall_firewall_policy" "pike" {
  name        = "pike"
  description = "pikey"
  firewall_policy {
    stateless_default_actions          = ["aws:pass"]
    stateless_fragment_default_actions = ["aws:drop"]
    stateless_rule_group_reference {
      priority     = 1
      resource_arn = aws_networkfirewall_rule_group.pike.arn
    }
  }

  tags = {
    pike = "permissions"
  }
}
