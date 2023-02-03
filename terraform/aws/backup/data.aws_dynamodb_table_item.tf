data "aws_dynamodb_table_item" "pike" {
  table_name = "pike"
  expression_attribute_names = {
    "#P" = "Percentile"
  }
  projection_expression = "#P"
  key                   = <<KEY
{
    "hashKey": {"S": "example"}
}
KEY
}
