data "google_artifact_registry_packages" "pike" {
  location      = "us-central1"
  repository_id = "pike"
}

output "google_artifact_registry_packages" {
  value = data.google_artifact_registry_packages.pike
}
