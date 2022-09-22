resource "aws_lambda_layer_version" "pike" {
  s3_bucket           = "testbucketineu-west2"
  s3_key              = "bin.zip"
  layer_name          = "pike"
  compatible_runtimes = ["go1.x"]
}

output "layer" {
  value = aws_lambda_layer_version.pike
}
