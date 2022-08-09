data "aws_s3_bucket" "selected" {
  bucket = "trails-680235478471"
}

output "bucket" {
  value = data.aws_s3_bucket.selected
}
