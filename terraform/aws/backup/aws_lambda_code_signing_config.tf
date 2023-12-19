resource "aws_lambda_code_signing_config" "pike" {
  allowed_publishers {
    signing_profile_version_arns = [
      aws_signer_signing_profile.pike.arn,
    ]
  }

  policies {
    untrusted_artifact_on_deployment = "Warn"
  }

  description = "My awesome code signing config."
}
