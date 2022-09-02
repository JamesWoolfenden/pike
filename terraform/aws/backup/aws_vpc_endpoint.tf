resource "aws_vpc_endpoint" "s3" {
  vpc_id       = "vpc-03036aea287d9ee9b"
  service_name = "com.amazonaws.eu-west-2.s3"

  tags = {
    pike = "permissions"
  }
}
