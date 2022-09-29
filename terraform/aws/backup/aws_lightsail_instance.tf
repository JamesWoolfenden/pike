resource "aws_lightsail_instance" "pike" {
  name              = "pike"
  availability_zone = "eu-west-2b"
  blueprint_id      = "amazon_linux"
  bundle_id         = "nano_2_0"
  key_pair_name     = aws_lightsail_key_pair.pike.name
  tags = {
    pike = "permissions"
    foo  = "bar"
  }
}
