data "aws_vpc_ipams" "pike" {
}

output "aws_vpc_ipams" {
  value = data.aws_vpc_ipams.pike
}
