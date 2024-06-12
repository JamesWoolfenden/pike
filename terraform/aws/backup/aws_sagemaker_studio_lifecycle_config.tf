resource "aws_sagemaker_studio_lifecycle_config" "pike" {
  studio_lifecycle_config_name     = "example"
  studio_lifecycle_config_app_type = "JupyterServer"
  studio_lifecycle_config_content  = base64encode("echo Hello")
  tags = {
    pike = "permissions"
  }
}
