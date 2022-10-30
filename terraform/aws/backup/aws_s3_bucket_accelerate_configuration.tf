resource "aws_s3_bucket_accelerate_configuration" "pike" {
  bucket = "pike-680235478471"
  status = "Enabled"
}
