resource "aws_sagemaker_endpoint_configuration" "pike" {
  name        = "pike-config"
  kms_key_arn = "arn:aws:kms:eu-west-2:680235478471:key/03a3077b-1b63-4f42-98a1-20ea7a2fabba"

  production_variants {
    variant_name           = "variant-1"
    model_name             = aws_sagemaker_model.pike.name
    initial_instance_count = 1
    instance_type          = "ml.t2.medium"
  }

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
