resource "aws_globalaccelerator_custom_routing_endpoint_group" "pike" {
  listener_arn = aws_globalaccelerator_custom_routing_listener.pike.id

  destination_configuration {
    from_port = 80
    to_port   = 8080
    protocols = ["TCP"]
  }

  endpoint_configuration {
    endpoint_id = data.aws_subnets.pike.ids[0]
  }
}

data "aws_subnets" "pike" {}
