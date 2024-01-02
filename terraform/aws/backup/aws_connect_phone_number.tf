resource "aws_connect_phone_number" "pike" {

  target_arn   = aws_connect_instance.pike.arn
  country_code = "US"
  type         = "DID"

  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
