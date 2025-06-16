data "google_parameter_manager_parameter_version_render" "pike" {
  parameter_version_id = "pike"
  parameter            = "pike"
}

output "google_parameter_manager_parameter_version_render" {
  value = data.google_parameter_manager_parameter_version_render.pike
}
