resource "aws_acmpca_certificate_authority" "pike" {
  certificate_authority_configuration {
    key_algorithm     = "RSA_4096"
    signing_algorithm = "SHA512WITHRSA"

    subject {
      common_name = "freebeer.site"
    }
  }

  permanent_deletion_time_in_days = 7
}
