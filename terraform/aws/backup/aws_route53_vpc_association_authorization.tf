resource "aws_route53_vpc_association_authorization" "pike" {
  vpc_id  = aws_vpc.alternate.id
  zone_id = aws_route53_zone.example.id
}
