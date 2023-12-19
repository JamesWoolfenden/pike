resource "aws_signer_signing_profile" "pike" {
  platform_id = "AWSLambda-SHA384-ECDSA"


  signature_validity_period {
    value = 135
    type  = "MONTHS"
  }

  tags = {
    pike = "permissions"
  }
}
