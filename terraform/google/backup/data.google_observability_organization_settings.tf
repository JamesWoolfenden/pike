data "google_observability_organization_settings" "pike" {
  provider = google-beta
}

output "google_observability_organization_settings" {
  value = data.google_observability_organization_settings.pike
}
