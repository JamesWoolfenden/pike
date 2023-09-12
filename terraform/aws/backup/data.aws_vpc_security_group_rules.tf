data "aws_vpc_security_group_rules" "pike" {}

output "rules" {
  value = data.aws_vpc_security_group_rules.pike
}
