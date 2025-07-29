data "aws_s3_access_point" "pike" {
  name = "pike"
}

output "aws_s3_access_point" {
  value = data.aws_s3_access_point.pike
}
