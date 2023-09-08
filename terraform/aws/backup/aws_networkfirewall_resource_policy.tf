resource "aws_networkfirewall_resource_policy" "pike" {
  resource_arn = aws_networkfirewall_firewall_policy.pike.arn
  # policy's Action element must include all of the following operations
  policy = jsonencode({
    Statement = [{
      Action = [
        "network-firewall:ListFirewallPolicies",
        "network-firewall:CreateFirewall",
        "network-firewall:UpdateFirewall",
        "network-firewall:AssociateFirewallPolicy"
      ]
      Effect   = "Allow"
      Resource = aws_networkfirewall_firewall_policy.pike.arn
      Principal = {
        AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      }
    }]
    Version = "2012-10-17"
  })

}

data "aws_caller_identity" "current" {}
