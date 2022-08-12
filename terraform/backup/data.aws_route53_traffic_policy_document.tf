data "aws_region" "current" {}

data "aws_route53_traffic_policy_document" "example" {
  record_type = "A"
  start_rule  = "site_switch"

  endpoint {
    id    = "my_elb"
    type  = "elastic-load-balancer"
    value = "elb-111111.${data.aws_region.current.name}.elb.amazonaws.com"
  }
  endpoint {
    id     = "site_down_banner"
    type   = "s3-website"
    region = data.aws_region.current.name
    value  = "www.example.com"
  }

  rule {
    id   = "site_switch"
    type = "failover"

    primary {
      endpoint_reference = "my_elb"
    }
    secondary {
      endpoint_reference = "site_down_banner"
    }
  }
}
