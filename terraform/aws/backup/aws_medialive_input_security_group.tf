resource "aws_medialive_input_security_group" "pike" {
  whitelist_rules {
    cidr = "10.0.0.8/32"
  }
  tags = {
    pike = "permissions"
  }
}
