resource "aws_datazone_environment_profile" "pike" {
  aws_account_id                   = data.aws_caller_identity.test.account_id
  aws_account_region               = data.aws_region.test.name
  description                      = "description"
  environment_blueprint_identifier = data.aws_datazone_environment_blueprint.test.id
  name                             = "example-name"
  project_identifier               = aws_datazone_project.test.id
  domain_identifier                = aws_datazone_domain.test.id
  user_parameters {
    name  = "consumerGlueDbName"
    value = "value"
  }
}

data "aws_caller_identity" "test" {}
data "aws_region" "test" {}
data "aws_datazone_environment_blueprint" "test" {
  domain_id = ""
  managed   = false
  name      = "pike"
}
