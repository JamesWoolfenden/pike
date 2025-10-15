data "aws_odb_network" "pike" {
  id = "pike"
}

output "aws_odb_network" {
  value = data.aws_odb_network.pike
}
