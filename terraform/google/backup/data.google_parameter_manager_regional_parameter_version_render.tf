data "google_parameter_manager_regional_parameter_version_render" "pike" {
  parameter_version_id = "pike"
  parameter            = "pike"
  location             = "us-central1"
}

output "google_parameter_manager_regional_parameter_version_render" {
  value = data.google_parameter_manager_regional_parameter_version_render.pike
}
