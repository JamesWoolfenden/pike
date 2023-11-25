data "aws_bedrock_foundation_model" "pike" {
  provider = aws.central
  model_id = "meta.llama2-13b-chat-v1"
}

output "model" {
  value = data.aws_bedrock_foundation_model.pike
}
