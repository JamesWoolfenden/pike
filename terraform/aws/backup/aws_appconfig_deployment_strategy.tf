resource "aws_appconfig_deployment_strategy" "pike" {
  name                           = "example-deployment-strategy-tf"
  description                    = "Example Deployment Strategy"
  deployment_duration_in_minutes = 3
  final_bake_time_in_minutes     = 4
  growth_factor                  = 10
  growth_type                    = "LINEAR"
  replicate_to                   = "NONE"

  tags = {
    Type = "AppConfig Deployment Strategy"
  }
}
