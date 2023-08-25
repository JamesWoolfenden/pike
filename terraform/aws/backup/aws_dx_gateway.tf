resource "aws_dx_gateway" "pike" {
  name            = "tf-dxg-example"
  amazon_side_asn = "64512"
}
