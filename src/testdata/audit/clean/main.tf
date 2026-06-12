resource "aws_iam_policy" "scoped" {
  name = "x"
  policy = jsonencode({
    Statement = [{
      Effect   = "Allow"
      Action   = ["s3:GetObject"]
      Resource = ["arn:aws:s3:::bucket/*"]
    }]
  })
}

resource "google_project_iam_member" "viewer" {
  project = "p"
  role    = "roles/viewer"
  member  = "user:a@example.com"
}

resource "azurerm_role_assignment" "reader" {
  scope                = "/subscriptions/x"
  role_definition_name = "Reader"
  principal_id         = "00000000-0000-0000-0000-000000000000"
}
