data "google_observability_folder_settings" "pike" {
  provider = google-beta
  location = "us-central1"
  folder   = "pike"
}

output "google_observability_folder_settings" {
  value = data.google_observability_folder_settings.pike
}
