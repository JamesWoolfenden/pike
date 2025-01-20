resource "aws_appconfig_deployment" "pike" {
  application_id           = aws_appconfig_application.example.id
  configuration_profile_id = aws_appconfig_configuration_profile.example.configuration_profile_id
  configuration_version    = aws_appconfig_hosted_configuration_version.example.version_number
  deployment_strategy_id   = aws_appconfig_deployment_strategy.example.id
  description              = "My example deployment"
  environment_id           = aws_appconfig_environment.example.environment_id
  kms_key_identifier       = aws_kms_key.example.arn

  tags = {
    Type = "AppConfig Deployment"
  }
}
