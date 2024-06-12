resource "aws_sagemaker_app" "pike" {
  domain_id         = aws_sagemaker_domain.pike.id
  user_profile_name = aws_sagemaker_user_profile.pike.user_profile_name
  app_name          = "example"
  app_type          = "JupyterServer"
  tags = {
    pike = "permissions"
  }
}
