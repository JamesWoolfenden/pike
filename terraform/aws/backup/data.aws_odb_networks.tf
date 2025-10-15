data "aws_odb_networks" "pike" {
}

output "aws_odb_networks" {
  value = data.aws_odb_networks.pike
}
