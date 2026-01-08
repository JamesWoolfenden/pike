resource "google_iam_oauth_client" "oauth_client" {
  oauth_client_id       = "example-client-id"
  location              = "global"
  allowed_grant_types   = ["AUTHORIZATION_CODE_GRANT"]
  allowed_redirect_uris = ["https://www.example.com"]
  allowed_scopes        = ["https://www.googleapis.com/auth/cloud-platform"]
  client_type           = "CONFIDENTIAL_CLIENT"
}
