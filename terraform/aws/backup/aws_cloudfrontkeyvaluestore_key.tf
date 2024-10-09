resource "aws_cloudfront_key_value_store" "example" {
  name    = "ExampleKeyValueStore"
  comment = "This is an example key value store"
}

resource "aws_cloudfrontkeyvaluestore_key" "example" {
  key_value_store_arn = aws_cloudfront_key_value_store.example.arn
  key                 = "Test Key"
  value               = "Test Value"
}
