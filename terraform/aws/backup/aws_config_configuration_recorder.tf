resource "aws_config_configuration_recorder" "pike" {
  name     = "pike"
  role_arn = "arn:aws:iam::680235478471:role/aws-config-config-680235478471"
}
