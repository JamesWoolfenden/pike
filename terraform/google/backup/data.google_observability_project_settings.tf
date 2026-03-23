data "google_observability_project_settings" "pike" {
  provider = google-beta
  project  = "pike"
  location = "us-central1"
}

output "google_observability_project_settings" {
  value = data.google_observability_project_settings.pike
}
