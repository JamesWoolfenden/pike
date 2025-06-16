data "google_parameter_manager_parameter" "pike" {
  parameter_id = "pike"
}

output "google_parameter_manager_parameter" {
  value = data.google_parameter_manager_parameter.pike
}
