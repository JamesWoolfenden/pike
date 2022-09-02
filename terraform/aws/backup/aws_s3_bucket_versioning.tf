resource "aws_s3_bucket_versioning" "example" {
  bucket = "pike-680235478471"
  versioning_configuration {
    status = "Enabled"
  }
  expected_bucket_owner = "680235478471"
}
