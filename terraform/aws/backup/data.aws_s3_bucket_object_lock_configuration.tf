data "aws_s3_bucket_object_lock_configuration" "pike" {
  bucket = "pike"
}

output "aws_s3_bucket_object_lock_configuration" {
  value = data.aws_s3_bucket_object_lock_configuration.pike
}
