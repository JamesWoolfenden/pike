data "google_artifact_registry_repositories" "pike" {
  location = "us-central1"
}

output "google_artifact_registry_repositories" {
  value = data.google_artifact_registry_repositories.pike
}
