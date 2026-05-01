data "google_logging_project_cmek_settings" "pike" {
  project = "pike-gcp"
}

output "google_logging_project_cmek_settings" {
  value = data.google_logging_project_cmek_settings.pike
}
