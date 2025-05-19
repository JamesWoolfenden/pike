data "aws_dynamodb_tables" "pike" {
}

output "aws_dynamodb_tables" {
  value = data.aws_dynamodb_tables.pike
}
