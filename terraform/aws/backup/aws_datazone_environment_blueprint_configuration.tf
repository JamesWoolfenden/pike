resource "aws_datazone_environment_blueprint_configuration" "pike" {
  domain_id                = aws_datazone_domain.pike.id
  environment_blueprint_id = data.aws_datazone_environment_blueprint.pike.id
  enabled_regions          = ["eu-west-2"]

  regional_parameters = {
    us-east-1 = {
      S3Location = "s3://my-amazon-datazone-bucket"
    }
  }
}
