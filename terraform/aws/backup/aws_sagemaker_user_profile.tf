resource "aws_sagemaker_user_profile" "pike" {
  domain_id         = aws_sagemaker_domain.pike.id
  user_profile_name = "james"
}
