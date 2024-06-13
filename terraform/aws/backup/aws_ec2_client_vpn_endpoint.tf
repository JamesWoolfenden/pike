resource "aws_ec2_client_vpn_endpoint" "pike" {
  description            = "terraform-clientvpn-example"
  server_certificate_arn = aws_acm_certificate.pike.arn
  client_cidr_block      = "10.0.0.0/16"

  authentication_options {
    type                       = "certificate-authentication"
    root_certificate_chain_arn = aws_acm_certificate.pike.arn
  }

  connection_log_options {
    enabled               = true
    cloudwatch_log_group  = aws_cloudwatch_log_group.lg.name
    cloudwatch_log_stream = aws_cloudwatch_log_stream.ls.name
  }
}

resource "aws_acm_certificate" "pike" {
  domain_name       = "pike.freebeer.site"
  validation_method = "EMAIL"

  tags = {
    pike = "permissions"
    #    Environment = "test"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_cloudwatch_log_group" "lg" {
  name = "pike"
}

resource "aws_cloudwatch_log_stream" "ls" {
  log_group_name = aws_cloudwatch_log_group.lg.name
  name           = "pike"
}
