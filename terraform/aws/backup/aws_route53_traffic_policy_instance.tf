resource "aws_route53_traffic_policy_instance" "pike" {
  name                   = "test.freebeer.site"
  traffic_policy_id      = aws_route53_traffic_policy.pike.id
  traffic_policy_version = 1
  hosted_zone_id         = data.aws_route53_zone.pike.zone_id
  ttl                    = 360
}

data "aws_route53_zone" "pike" {
  zone_id = "Z0613304D03LG1SU5BI"
}
