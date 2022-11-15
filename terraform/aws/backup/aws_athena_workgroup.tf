resource "aws_athena_workgroup" "pike" {
  name = "pike"

  configuration {
    enforce_workgroup_configuration    = true
    publish_cloudwatch_metrics_enabled = true

    result_configuration {
      output_location = "s3://athenaresults-mp42nr/output/"

      encryption_configuration {
        encryption_option = "SSE_KMS"
        kms_key_arn       = "arn:aws:kms:eu-west-2:680235478471:key/8d2c4754-28b5-4590-8627-e4ac0792980e"
      }
    }
  }
  tags = {
    pike = "permissions"
  }
}
