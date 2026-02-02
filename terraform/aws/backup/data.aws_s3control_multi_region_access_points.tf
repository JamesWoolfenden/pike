data "aws_s3control_multi_region_access_points" "pike" {
  provider = aws.uswest2
}

output "aws_s3control_multi_region_access_points" {
  value = data.aws_s3control_multi_region_access_points.pike
}
