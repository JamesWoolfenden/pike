data "aws_secretsmanager_secrets" "pike" {}

output "sec" {
  value     = data.aws_secretsmanager_secrets.pike
  sensitive = true
}
