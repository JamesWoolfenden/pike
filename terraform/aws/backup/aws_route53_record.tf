resource "aws_route53_record" "example" {
  name            = "example"
  type            = "A"
  ttl             = 100
  zone_id         = "Z0613304D03LG1SU5BI"
  records         = ["8.8.8.8"]
  allow_overwrite = true
}
