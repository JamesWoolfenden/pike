data "aws_vpclattice_listener" "pike" {
  listener_identifier = "arn:aws:vpc-lattice:eu-west-2:680235478471:service/svc-12313212312313345/listener/listener-12341231323123123"
  service_identifier  = "arn:aws:vpc-lattice:eu-west-2:680235478471:servicenetwork/sn-12341231323123123"
}
