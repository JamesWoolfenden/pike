resource "google_artifact_registry_project_config" "pike" {
  location = "us-central1"
  platform_logs_config {
    logging_state = "ENABLED"
  }
}
