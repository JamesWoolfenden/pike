resource "aws_cloudwatch_event_endpoint" "pike" {
  name     = "global-endpoint"
  role_arn = "arn:aws:iam::680235478471:role/aws-service-role/apidestinations.events.amazonaws.com/AWSServiceRoleForAmazonEventBridgeApiDestinations"

  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.primary.arn
  }
  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.secondary.arn
  }

  replication_config {
    state = "DISABLED"
  }

  routing_config {
    failover_config {
      primary {
        health_check = aws_route53_health_check.primary.arn
      }

      secondary {
        route = "us-east-2"
      }
    }
  }

}

resource "aws_cloudwatch_event_bus" "primary" {
  name = "chat-messages"
}

resource "aws_cloudwatch_event_bus" "secondary" {
  name = "guff-messages"
}

resource "aws_route53_health_check" "primary" {
  fqdn              = "example.com"
  port              = 80
  type              = "HTTP"
  resource_path     = "/"
  failure_threshold = "5"
  request_interval  = "30"

  tags = {
    Name = "tf-test-health-check"
  }
}
