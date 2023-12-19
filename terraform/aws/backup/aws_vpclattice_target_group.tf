resource "aws_vpclattice_target_group" "pike" {
  name = "example"
  type = "INSTANCE"

  config {
    vpc_identifier = "vpc-0c33dc8cd64f408c4"

    port     = 443
    protocol = "HTTPS"
  }

}
