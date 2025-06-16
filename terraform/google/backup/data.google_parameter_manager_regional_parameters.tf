data "google_parameter_manager_regional_parameters" "pike" {
  location = "us-central1"
}

output "google_parameter_manager_regional_parameters" {
  value = data.google_parameter_manager_regional_parameters.pike
}
