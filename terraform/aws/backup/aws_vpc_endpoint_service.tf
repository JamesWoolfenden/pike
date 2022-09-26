resource "aws_vpc_endpoint_service" "pike" {
  acceptance_required        = false
  network_load_balancer_arns = [aws_lb.net.arn]
  // gateway_load_balancer_arns = []
  tags = {
    pike = "permissions"
  }
}

resource "aws_lb" "net" {
  name               = "test-lb-tf"
  internal           = false
  load_balancer_type = "network"
  subnets            = data.aws_subnet_ids.public.ids

  enable_deletion_protection = false

  tags = {
    Environment = "production"
  }
}

data "aws_subnet_ids" "public" {
  vpc_id = "vpc-0c33dc8cd64f408c4"
}

output "subnets" {
  value = data.aws_subnet_ids.public
}
