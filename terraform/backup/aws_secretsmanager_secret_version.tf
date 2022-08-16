resource "aws_secretsmanager_secret_version" "pike" {
  secret_id     = "arn:aws:secretsmanager:eu-west-2:680235478471:secret:pike-2kiLPh"
  secret_string = "example-string-to-protect"
}
