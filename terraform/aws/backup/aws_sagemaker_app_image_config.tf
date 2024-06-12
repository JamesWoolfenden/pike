resource "aws_sagemaker_app_image_config" "pike" {
  app_image_config_name = "example"

  kernel_gateway_image_config {
    kernel_spec {
      name = "example"
    }
  }
  tags = {
    pike = "permissions"
  }
}
