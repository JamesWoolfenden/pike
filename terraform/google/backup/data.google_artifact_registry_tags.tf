data "google_artifact_registry_tags" "pike" {
  location      = "us-central1"
  package_name  = "pike"
  repository_id = "pike"
}

output "google_artifact_registry_tags" {
  value = data.google_artifact_registry_tags.pike
}
