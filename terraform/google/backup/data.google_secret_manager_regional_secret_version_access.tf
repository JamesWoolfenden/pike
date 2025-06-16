data "google_secret_manager_regional_secret_version_access" "pike" {
  secret   = "my-secret"
  location = "us-central1"
}

output "google_secret_manager_regional_secret_version_access" {
  value = data.google_secret_manager_regional_secret_version_access.pike
}
