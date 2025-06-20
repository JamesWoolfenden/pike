resource "aws_lightsail_bucket_access_key" "pike" {
  bucket_name = aws_lightsail_bucket.pike.name
}

resource "aws_lightsail_bucket" "pike" {
  bundle_id = "small_1_0"
  name      = "jgw-lightsail-bucket2"
}
