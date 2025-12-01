data "google_artifact_registry_python_packages" "pike" {
  location      = "europe-west1"
  repository_id = "pike"
}

output "google_artifact_registry_python_packages" {
  value = data.google_artifact_registry_python_packages.pike
}
