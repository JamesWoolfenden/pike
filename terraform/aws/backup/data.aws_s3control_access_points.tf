data "aws_s3control_access_points" "pike" {
}

output "aws_s3control_access_points" {
  value = data.aws_s3control_access_points.pike
}
