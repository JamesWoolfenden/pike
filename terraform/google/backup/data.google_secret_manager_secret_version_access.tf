data "google_secret_manager_secret_version_access" "pike" {
  secret = "mysecret"
}
