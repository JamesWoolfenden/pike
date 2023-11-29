resource "aws_cloudsearch_domain" "pike" {
  provider = aws.central
  name     = "freebeer"
}
