data "google_storage_control_project_intelligence_config" "pike" {
  name = "pike-477416"
}

output "google_storage_control_project_intelligence_config" {
  value = data.google_storage_control_project_intelligence_config.pike
}
