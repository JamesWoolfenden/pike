
data "aws_s3_bucket_object" "example" {
  bucket = "testbucketforlbjgw"
  key    = "prefix/AWSLogs/680235478471/ELBAccessLogTestFile"
}

output "bucket_object" {
  value = data.aws_s3_bucket_object.example
}

data "aws_s3_object" "example" {
  bucket = "testbucketforlbjgw"
  key    = "prefix/AWSLogs/680235478471/ELBAccessLogTestFile"
}

output "object" {
  value = data.aws_s3_object.example
}
