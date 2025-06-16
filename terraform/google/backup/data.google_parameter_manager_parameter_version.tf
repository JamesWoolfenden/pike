data "google_parameter_manager_parameter_version" "pike" {
  parameter_version_id = "pike"
  parameter            = "pike"
}

output "google_parameter_manager_parameter_version" {
  value = data.google_parameter_manager_parameter_version.pike
}
