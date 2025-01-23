resource "aws_globalaccelerator_custom_routing_listener" "pike" {
  accelerator_arn = aws_globalaccelerator_custom_routing_accelerator.pike.id

  port_range {
    from_port = 80
    to_port   = 80
  }
}
