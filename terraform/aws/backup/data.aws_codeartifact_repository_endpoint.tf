data "aws_codeartifact_repository_endpoint" "pike" {
  domain     = "pike"
  format     = "npm"
  repository = "pike"
}
