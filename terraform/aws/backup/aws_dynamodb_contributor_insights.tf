resource "aws_dynamodb_contributor_insights" "pike" {
  table_name = data.aws_dynamodb_table.pike.name
}
