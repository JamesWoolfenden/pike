data "google_artifact_registry_locations" "pike" {
}

output "google_artifact_registry_locations" {
  value = data.google_artifact_registry_locations.pike
}
