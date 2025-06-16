data "google_parameter_manager_regional_parameter" "pike" {
  location     = "us-central1"
  parameter_id = "pike"
}

output "google_parameter_manager_regional_parameter" {
  value = data.google_parameter_manager_regional_parameter.pike
}
