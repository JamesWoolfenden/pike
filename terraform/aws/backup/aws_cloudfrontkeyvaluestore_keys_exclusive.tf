resource "aws_cloudfront_key_value_store" "pike" {
  name    = "ExampleKeyValueStore"
  comment = "This is an example key value store"
}

resource "aws_cloudfrontkeyvaluestore_keys_exclusive" "pike" {
  key_value_store_arn = aws_cloudfront_key_value_store.pike.arn

  resource_key_value_pair {
    key   = "Test Key"
    value = "Test Value"
  }
}
