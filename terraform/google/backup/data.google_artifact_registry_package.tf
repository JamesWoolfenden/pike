data "google_artifact_registry_package" "pike" {
  location      = "us-central1"
  name          = "pike"
  repository_id = "pike"
}

output "google_artifact_registry_package" {
  value = data.google_artifact_registry_package.pike
}
