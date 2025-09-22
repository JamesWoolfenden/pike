data "google_artifact_registry_python_package" "pike" {
  package_name  = "pike"
  repository_id = "pike"
  location      = "us-central1"
}

output "google_artifact_registry_python_package" {
  value = data.google_artifact_registry_python_package.pike
}
