data "google_secret_manager_regional_secrets" "pike" {
  location = "us-central1"
}

output "google_secret_manager_regional_secrets" {
  value = data.google_secret_manager_regional_secrets.pike
}
