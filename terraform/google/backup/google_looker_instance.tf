resource "google_looker_instance" "pike" {
  name             = "my-instance"
  platform_edition = "LOOKER_CORE_STANDARD_ANNUAL"
  region           = "us-central1"
  oauth_config {
    client_id     = "my-client-id"
    client_secret = "my-client-secret"
  }
  deletion_policy = "DEFAULT"
}
