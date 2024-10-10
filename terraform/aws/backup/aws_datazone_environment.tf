resource "aws_datazone_environment" "pike" {
  name                 = "example"
  account_identifier   = data.aws_caller_identity.test.account_id
  account_region       = data.aws_region.test.name
  blueprint_identifier = aws_datazone_environment_blueprint_configuration.test.environment_blueprint_id
  profile_identifier   = aws_datazone_environment_profile.test.id
  project_identifier   = aws_datazone_project.test.id
  domain_identifier    = aws_datazone_domain.test.id

  user_parameters {
    name  = "consumerGlueDbName"
    value = "consumer"
  }

  user_parameters {
    name  = "producerGlueDbName"
    value = "producer"
  }

  user_parameters {
    name  = "workgroupName"
    value = "workgroup"
  }
}
