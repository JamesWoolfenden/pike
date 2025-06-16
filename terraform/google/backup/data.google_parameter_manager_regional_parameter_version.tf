data "google_parameter_manager_regional_parameter_version" "pike" {
  parameter_version_id = "pike"
  parameter            = "pike"
  location             = "us-central1"
}

output "google_parameter_manager_regional_parameter_version" {
  value = data.google_parameter_manager_regional_parameter_version.pike
}
