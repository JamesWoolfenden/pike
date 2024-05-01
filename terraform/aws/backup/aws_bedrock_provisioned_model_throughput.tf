resource "aws_bedrock_provisioned_model_throughput" "pike" {
  provider               = aws.central
  provisioned_model_name = "example-model"
  model_arn              = aws_bedrock_custom_model.pike.custom_model_arn
  commitment_duration    = "SixMonths"
  model_units            = 1
}
