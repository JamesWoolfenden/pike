data "aws_eip" "pike" {}

output "ip" {
  value = data.aws_eip.pike
}
