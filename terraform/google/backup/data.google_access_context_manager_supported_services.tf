data "google_access_context_manager_supported_services" "pike" {
}

output "google_access_context_manager_supported_services" {
  value = data.google_access_context_manager_supported_services.pike
}
