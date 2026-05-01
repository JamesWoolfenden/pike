data "google_secret_manager_secrets" "pike" {
  provider = google-beta
}

output "google_secret_manager_secrets" {
  value = data.google_secret_manager_secrets.pike
}
