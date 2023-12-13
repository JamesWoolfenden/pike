data "google_secret_manager_secrets" "pike" {
  provider = google-beta
}

output "secrets" {
  value = data.google_secret_manager_secrets.pike
}
