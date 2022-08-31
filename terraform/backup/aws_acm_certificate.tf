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
