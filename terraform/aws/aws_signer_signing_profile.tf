resource "aws_signer_signing_profile" "pike" {
  platform_id = "AWSLambda-SHA384-ECDSA"
  name_prefix = "prod_sp_"

  signature_validity_period {
    value = 5
    type  = "YEARS"
  }

  tags = {
    tag1 = "value1"
    tag2 = "value2"
  }
}