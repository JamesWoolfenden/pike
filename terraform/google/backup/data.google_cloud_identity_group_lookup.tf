data "google_cloud_identity_group_lookup" "pike" {
  provider = google-beta
  group_key {
    id = "my-group@example.com"
  }
}
