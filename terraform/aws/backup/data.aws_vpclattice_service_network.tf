data "aws_vpclattice_service_network" "pike" {
  service_network_identifier = "arn:aws:vpc-lattice:eu-west-2:680235478471:servicenetwork/sn-12341231323123123"
}

output "network" {
  value = data.aws_vpclattice_service_network.pike
}
