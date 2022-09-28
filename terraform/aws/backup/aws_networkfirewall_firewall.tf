resource "aws_networkfirewall_firewall" "pike" {
  description                       = "change"
  firewall_policy_arn               = aws_networkfirewall_firewall_policy.pike.arn
  firewall_policy_change_protection = false
  subnet_change_protection          = false

  name   = "pike"
  vpc_id = "vpc-00ea5eff890b0e212"
  subnet_mapping {
    subnet_id = "subnet-043bb893867355740"
  }
  subnet_mapping {
    subnet_id = "subnet-05e87623a2a5c41f0"
  }
  subnet_mapping {
    subnet_id = "subnet-09ff91b5b0adb1fd4"
  }
  delete_protection = false

  tags = {
    pike = "permissions"
  }
}
