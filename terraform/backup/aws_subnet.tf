resource "aws_subnet" "example" {
  vpc_id     = "vpc-0c9622709bb598517"
  cidr_block = "10.0.0.0/24"

  #   tags = {
  #     Name = "Main"
  #   }
}
