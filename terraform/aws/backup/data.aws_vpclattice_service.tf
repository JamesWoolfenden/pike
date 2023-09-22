data "aws_vpclattice_service" "pike" {
  name = "pike"
}

output "service" {
  value = data.aws_vpclattice_service.pike
}
