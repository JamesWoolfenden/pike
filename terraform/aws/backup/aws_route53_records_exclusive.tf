resource "aws_route53_records_exclusive" "pike" {
  zone_id = aws_route53_zone.test.zone_id

  resource_record_set {
    name = "subdomain.freebeer.site"
    type = "A"
    ttl  = "30"

    resource_records {
      value = "127.0.0.1"
    }
    resource_records {
      value = "127.0.0.30"
    }
  }
}


resource "aws_route53_zone" "test" {
  name          = "freebeer.site"
  force_destroy = true
}
