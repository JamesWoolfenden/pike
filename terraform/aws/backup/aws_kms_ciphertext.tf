resource "aws_kms_ciphertext" "pike" {
  provider = aws.central
  key_id   = "arn:aws:kms:us-east-1:680235478471:key/d983005c-89e3-439b-b999-b6b587d3a3a7"

  plaintext = <<EOF
{
  "client_id": "e000dbae00000f00da00",
  "client_secret": "0000000d00000ace00e0000ec00000000000b0a0"
}
EOF
}
