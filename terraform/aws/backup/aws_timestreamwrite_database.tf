resource "aws_timestreamwrite_database" "pike" {
  provider      = aws.central
  database_name = "database-example"
  kms_key_id    = "arn:aws:kms:us-east-1:680235478471:key/d983005c-89e3-439b-b999-b6b587d3a3a7"

  tags = {
    pike = "permissions"
    #    Name = "value"
  }

}
