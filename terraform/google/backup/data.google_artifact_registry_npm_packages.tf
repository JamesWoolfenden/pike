data "google_artifact_registry_npm_packages" "pike" {
  repository_id = "pike"
  location      = "us-central1"
}

output "google_artifact_registry_npm_packages" {
  value = data.google_artifact_registry_npm_packages.pike
}
