resource "aws_route53_zone_association" "example" {
  zone_id = "example"
  vpc_id  = "example"
}
