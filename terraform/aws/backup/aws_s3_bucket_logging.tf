resource "aws_s3_bucket_logging" "example" {
  bucket        = "testbucketineu-west2"
  target_bucket = "logging-680235478471"
  target_prefix = "log/"
}
