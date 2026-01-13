resource "google_artifact_registry_package" "pike" {
  location      = "us-west1"
  repository_id = "my-repository"
}
