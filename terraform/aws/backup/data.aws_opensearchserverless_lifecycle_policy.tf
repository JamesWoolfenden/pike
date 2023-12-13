data "aws_opensearchserverless_lifecycle_policy" "pike" {
  provider = aws.central
  name     = "example-lifecycle-policy"
  type     = "retention"
}
