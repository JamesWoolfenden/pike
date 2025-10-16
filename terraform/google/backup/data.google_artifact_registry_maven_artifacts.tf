data "google_artifact_registry_maven_artifacts" "pike" {
  repository_id = "pike"
  location      = "us-central1"
}

output "google_artifact_registry_maven_artifacts" {
  value = data.google_artifact_registry_maven_artifacts.pike
}
