data "aws_vpc_ipam" "pike" {
  id = "ipam-abcd1234"
}

output "aws_vpc_ipam" {
  value = data.aws_vpc_ipam.pike
}
