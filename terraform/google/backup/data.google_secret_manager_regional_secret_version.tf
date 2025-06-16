data "google_secret_manager_regional_secret_version" "pike" {
  secret   = "my-secret"
  location = "us-central1"
}

output "google_secret_manager_regional_secret_version" {
  value = data.google_secret_manager_regional_secret_version.pike
}
