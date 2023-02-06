resource "aws_dynamodb_table_item" "pike" {
table_name = "pike"
hash_key   = data.aws_dynamodb_table.pike.hash_key

item = <<ITEM
{
  "first":{"S":"test"},
  "exampleHashKey": {"S": "something"},
  "one": {"N": "11111"},
  "two": {"N": "22222"},
  "three": {"N": "33333"},
  "four": {"N": "44444"}
}
ITEM
}