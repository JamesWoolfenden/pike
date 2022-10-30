resource "aws_s3_bucket_ownership_controls" "pike" {
  bucket = "pike-680235478471"

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}
