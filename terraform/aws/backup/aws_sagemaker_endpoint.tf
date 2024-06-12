resource "aws_sagemaker_endpoint" "pike" {
  endpoint_config_name = "pike"
  tags = {
    pike = "permissions"
  }
}
