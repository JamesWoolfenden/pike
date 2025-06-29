data "google_storage_control_project_intelligence_config" "pike" {
  name = "pike-412922"
}

output "google_storage_control_project_intelligence_config" {
  value = data.google_storage_control_project_intelligence_config.pike
}
