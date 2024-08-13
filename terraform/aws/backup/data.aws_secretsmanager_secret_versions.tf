data "aws_secretsmanager_secret_versions" "pike" {
  secret_id = "pike"
}

output "aws_secretsmanager_secret_versions" {
  value = data.aws_secretsmanager_secret_versions.pike
}
