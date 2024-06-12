resource "aws_sagemaker_model_package_group" "pike" {
  model_package_group_name = "pike"
  tags = {
    pike = "permissions"
  }
}
