data "aws_bedrock_custom_models" "pike" {
  provider = aws.central
}

output "models" {
  value = data.aws_bedrock_custom_models.pike
}
