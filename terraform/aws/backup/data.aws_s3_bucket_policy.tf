data "aws_s3_bucket_policy" "pike" {
  bucket = "pike-680235478471"
}

output "policy" {
  value = data.aws_s3_bucket_policy.pike
}
