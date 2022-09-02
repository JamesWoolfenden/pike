

resource "aws_s3_bucket_server_side_encryption_configuration" "example" {
  bucket = "pike-680235478471"
  rule {
    apply_server_side_encryption_by_default {
      //kms_master_key_id = aws_kms_key.mykey.arn
      sse_algorithm = "aws:kms"
    }
  }
}
