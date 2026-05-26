data "aws_dynamodb_backups" "pike" {
}

output "aws_dynamodb_backups" {
  value = data.aws_dynamodb_backups.pike
}
