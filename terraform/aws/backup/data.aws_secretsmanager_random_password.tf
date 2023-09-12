data "aws_secretsmanager_random_password" "pike" {}

output "pass" {
  value     = data.aws_secretsmanager_random_password.pike
  sensitive = true
}
