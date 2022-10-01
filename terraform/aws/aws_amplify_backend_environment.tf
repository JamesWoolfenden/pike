resource "aws_amplify_backend_environment" "example" {
  app_id           = aws_amplify_app.pike.id
  environment_name = "pike"

  deployment_artifacts = "app-example-deployment"
  stack_name           = "amplify-app-example"
}
