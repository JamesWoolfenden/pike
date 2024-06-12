resource "aws_sagemaker_pipeline" "pike" {
  pipeline_display_name = "pike"
  pipeline_name         = "pike"

  pipeline_definition = jsonencode({
    Version = "2020-12-01"
    Steps = [{
      Name = "Test"
      Type = "Fail"
      Arguments = {
        ErrorMessage = "test"
      }
    }]
  })

  tags = {
    pike = "permissions"
  }
}
