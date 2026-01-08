
resource "google_iam_oauth_client_credential" "example" {
  oauthclient                = google_iam_oauth_client.oauth_client.oauth_client_id
  location                   = google_iam_oauth_client.oauth_client.location
  oauth_client_credential_id = "cred-id"
  disabled                   = true
  display_name               = "Display Name of credential"
}
