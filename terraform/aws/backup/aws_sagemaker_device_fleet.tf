resource "aws_sagemaker_device_fleet" "pike" {
  device_fleet_name = "pike"
  role_arn          = aws_iam_role.example.arn

  output_config {
    s3_output_location = "s3://${aws_s3_bucket.pike.bucket}/ prefix/"
  }
}

resource "aws_s3_bucket" "pike" {
  bucket = "pike123456789"
}
