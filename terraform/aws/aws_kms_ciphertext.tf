resource "aws_kms_ciphertext" "pike" {
  provider = aws.central
    key_id = "arn:aws:kms:us-east-1:680235478471:key/d983005c-89e3-439b-b999-b6b587d3a3a7"

    plaintext = <<EOF
{
  "client_id": "e587dbae22222f55da22",
  "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
}
EOF
  }
