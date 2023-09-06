data "aws_security_groups" "pike" {}

output "sg" {
  value = data.aws_security_groups.pike
}
