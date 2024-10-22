data "aws_bedrock_inference_profile" "pike" {
  inference_profile_id = "pike"
}

output "aws_bedrock_inference_profile" {
  value = data.aws_bedrock_inference_profile.pike
}
