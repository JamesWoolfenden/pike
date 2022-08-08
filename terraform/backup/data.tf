data "aws_vpcs" "main" {
  tags = {
    Name = "test"
  }
}

output "aws_vpcs" {
  value = data.aws_vpcs.main
}


data "aws_caller_identity" "current" {}

output "id" {
  value = data.aws_caller_identity.current
}

data "aws_subnet_ids" "public" {
  vpc_id = "vpc-0c33dc8cd64f408c4"
}

output "subnets" {
  value = data.aws_subnet_ids.public
}

data "aws_vpc" "this" {
  id = "vpc-0c33dc8cd64f408c4"
}
