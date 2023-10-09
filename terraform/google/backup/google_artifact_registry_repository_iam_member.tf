resource "google_artifact_registry_repository_iam_member" "pike" {
  project    = google_artifact_registry_repository.pike.project
  location   = google_artifact_registry_repository.pike.location
  repository = google_artifact_registry_repository.pike.name
  role       = "roles/artifactregistry.reader"
  member     = "user:anonymous@gmail.com"
}
