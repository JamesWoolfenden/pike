resource "aws_dynamodb_tag" "pike" {
  key          = "pike"
  resource_arn = data.aws_dynamodb_table.pike.arn
  value        = "permissions"
}
