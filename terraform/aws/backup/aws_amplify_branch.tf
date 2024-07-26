resource "aws_amplify_branch" "pike" {
  app_id = aws_amplify_app.pike.id

  framework = "React"
  stage     = "PRODUCTION"

  environment_variables = {
    REACT_APP_API_SERVER = "https://api.example.com"
  }
  branch_name = "my-branch"
}
