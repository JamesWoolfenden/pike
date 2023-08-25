resource "aws_acmpca_certificate_authority_certificate" "example" {
  certificate_authority_arn = data.aws_acmpca_certificate_authority.pike.arn

  certificate       = aws_acmpca_certificate.example.certificate
  certificate_chain = aws_acmpca_certificate.example.certificate_chain
}

resource "aws_acmpca_certificate" "example" {
  certificate_authority_arn   = data.aws_acmpca_certificate_authority.pike.arn
  certificate_signing_request = data.aws_acmpca_certificate_authority.pike.certificate_signing_request
  signing_algorithm           = "SHA512WITHRSA"

  template_arn = "arn:${data.aws_partition.current.partition}:acm-pca:::template/RootCACertificate/V1"

  validity {
    type  = "YEARS"
    value = 1
  }
}

resource "aws_acmpca_certificate_authority" "example" {
  type = "ROOT"

  certificate_authority_configuration {
    key_algorithm     = "RSA_4096"
    signing_algorithm = "SHA512WITHRSA"

    subject {
      common_name = "example.com"
    }
  }
  tags = {
    pike = "permissions"
  }
}

data "aws_partition" "current" {}
