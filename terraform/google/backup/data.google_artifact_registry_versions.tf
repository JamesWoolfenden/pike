data "google_artifact_registry_version" "pike" {
  location      = "us-central1"
  repository_id = "pike"
  package_name  = "pike"
  version_name  = "pike"
}

output "google_artifact_registry_version" {
  value = data.google_artifact_registry_version.pike
}
