resource "aws_s3_bucket_object_lock_configuration" "pike" {
  bucket = "pike-680235478471"

  rule {
    default_retention {
      mode = "COMPLIANCE"
      days = 5
    }
  }

  token = "sometoken=="
}
