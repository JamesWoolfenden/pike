data "aws_bedrock_foundation_models" "pike" {
  provider = aws.central
}

output "models" {
  value = data.aws_bedrock_foundation_models.pike
}
