resource "aws_verifiedaccess_endpoint" "pike" {
  application_domain     = "example.com"
  attachment_type        = "vpc"
  description            = "example"
  domain_certificate_arn = aws_acm_certificate.pike.arn
  endpoint_domain_prefix = "example"
  endpoint_type          = "network-interface"
  network_interface_options {
    network_interface_id = aws_network_interface.pike.id
    port                 = 443
    protocol             = "https"
  }
  security_group_ids       = [aws_security_group.pike.id]
  verified_access_group_id = aws_verifiedaccess_group.pike.id
}

resource "aws_security_group" "pike" {
}

resource "aws_network_interface" "pike" {
  subnet_id      = "subnet-0562ef1d304b968f4"
  description    = "An interface"
  interface_type = "trunk"
  #  ipv4_prefix_count = 1
  tags = {
    pike = "permissions"
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
