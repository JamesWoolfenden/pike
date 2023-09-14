resource "google_artifact_registry_repository_iam_binding" "pike" {
  project    = google_artifact_registry_repository.pike.project
  location   = google_artifact_registry_repository.pike.location
  repository = google_artifact_registry_repository.pike.name
  role       = "roles/artifactregistry.reader"
  members = [
    "user:james.woolfenden@gmail.com",
    "user:crwoolfenden@gmail.com"
  ]
}
