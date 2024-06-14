resource "aws_simpledb_domain" "pike" {
  provider = aws.central
  name     = "pike"
}
