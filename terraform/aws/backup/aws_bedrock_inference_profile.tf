resource "aws_bedrock_inference_profile" "pike" {
  name        = "Claude Sonnet for Project 123"
  description = "Profile with tag for cost allocation tracking"

  model_source {
    copy_from = "arn:aws:bedrock:us-west-2::foundation-model/anthropic.claude-3-5-sonnet-20241022-v2:0"

    # Include account ID to use inference profiles
    # copy_from = "arn:aws:bedrock:eu-central-1:${data.aws_caller_identity.current.account_id}:inference-profile/eu.anthropic.claude-3-5-sonnet-20240620-v1:0"
  }

  tags = {
    ProjectID = "123"
  }
}

data "aws_caller_identity" "current" {}
