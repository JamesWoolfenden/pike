data "google_parameter_manager_parameters" "pike" {
}

output "google_parameter_manager_parameters" {
  value = data.google_parameter_manager_parameters.pike
}
