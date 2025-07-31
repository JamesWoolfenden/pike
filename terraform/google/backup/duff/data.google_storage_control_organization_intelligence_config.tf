data "google_storage_control_organization_intelligence_config" "pike" {
  name = "pike"
}

output "google_storage_control_organization_intelligence_config" {
  value = data.google_storage_control_organization_intelligence_config.pike
}
