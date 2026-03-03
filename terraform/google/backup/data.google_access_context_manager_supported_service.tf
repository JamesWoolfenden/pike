data "google_access_context_manager_supported_service" "pike" {
  service_name = "pike"
}

output "google_access_context_manager_supported_service" {
  value = data.google_access_context_manager_supported_service.pike
}
