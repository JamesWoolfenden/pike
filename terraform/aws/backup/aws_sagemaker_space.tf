resource "aws_sagemaker_space" "pike" {
  domain_id  = aws_sagemaker_domain.pike.id
  space_name = "pike"
}
