resource "google_secret_manager_secret" "pike" {
  secret_id = "secret-version"

  labels = {
    label = "my-label"
  }

  replication {
    auto {}
  }
}
