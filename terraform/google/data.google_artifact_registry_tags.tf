data "google_artifact_registry_tags" "pike" {
}

output "google_artifact_registry_tags" {
  value = data.google_artifact_registry_tags.pike
}
