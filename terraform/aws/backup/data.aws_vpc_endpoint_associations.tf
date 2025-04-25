data "aws_vpc_endpoint_associations" "pike" {
  vpc_endpoint_id = "lskdfhaksjdfb"
}

output "aws_vpc_endpoint_associations" {
  value = data.aws_vpc_endpoint_associations.pike
}
