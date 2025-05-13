data "aws_lambda_layer_version" "pike" {
  layer_name = "pike"
}

output "aws_lambda_layer_version" {
  value = data.aws_lambda_layer_version.pike
}
