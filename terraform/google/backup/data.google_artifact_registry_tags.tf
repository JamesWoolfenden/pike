data "google_artifact_registry_tags" "pike" {
  repository_id = "pike"
  package_name  = "pike"
  location      = "us-central1"
}

output "google_artifact_registry_tags" {
  value = data.google_artifact_registry_tags.pike
}
