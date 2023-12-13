resource "aws_route53_cidr_location" "pike" {
  cidr_collection_id = aws_route53_cidr_collection.pike.id
  name               = "office"
  cidr_blocks        = ["200.5.3.0/24", "200.6.3.0/24"]
}
