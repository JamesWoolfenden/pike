data "aws_bedrock_inference_profiles" "pike" {
}

output "aws_bedrock_inference_profiles" {
  value = data.aws_bedrock_inference_profiles.pike
}
