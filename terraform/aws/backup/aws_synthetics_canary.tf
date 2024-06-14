resource "aws_synthetics_canary" "pike" {
  name = "pike"
  schedule {
    expression = ""
  }

  artifact_s3_location = ""
  handler              = ""
  runtime_version      = ""
  execution_role_arn   = ""
  tags = {
    pike = "permissions"
  }
}
