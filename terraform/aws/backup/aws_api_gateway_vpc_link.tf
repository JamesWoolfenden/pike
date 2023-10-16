resource "aws_lb" "example" {
  name               = "example"
  internal           = true
  load_balancer_type = "network"

  subnet_mapping {
    subnet_id = "subnet-0562ef1d304b968f4"
  }
}

resource "aws_api_gateway_vpc_link" "pike" {
  name        = "example"
  description = "example description"
  target_arns = [aws_lb.example.arn]
  tags = {
    pike = "permission"
  }
}
