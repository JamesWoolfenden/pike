resource "aws_iam_organizations_features" "pike" {
  enabled_features = [
    "RootCredentialsManagement",
    "RootSessions"
  ]
}
