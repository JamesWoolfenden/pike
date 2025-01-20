resource "aws_appsync_graphql_api" "pike" {
  authentication_type = "AWS_IAM"
  name                = "example"
}
