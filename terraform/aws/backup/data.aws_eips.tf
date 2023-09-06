data "aws_eips" "pike" {}

output "ips" {
  value = data.aws_eips.pike
}
