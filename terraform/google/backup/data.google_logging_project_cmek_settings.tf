data "google_logging_project_cmek_settings" "pike" {
  project = "pike-gcp"
}

output "cmek" {
  value = data.google_logging_project_cmek_settings.pike
}
