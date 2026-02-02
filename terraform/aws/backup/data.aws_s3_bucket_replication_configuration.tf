data "aws_s3_bucket_replication_configuration" "pike" {
  bucket = "pike"
}


output "aws_s3_bucket_replication_configuration" {
  value = data.aws_s3_bucket_replication_configuration.pike
}
