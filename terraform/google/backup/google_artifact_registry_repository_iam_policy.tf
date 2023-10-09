data "google_iam_policy" "admin" {
  binding {
    role = "roles/artifactregistry.reader"
    members = [
      "user:james.woolfenden@gmail.com",
      "user:anonymous@gmail.com",
    ]
  }
}

resource "google_artifact_registry_repository_iam_policy" "pike" {
  project     = google_artifact_registry_repository.pike.project
  location    = google_artifact_registry_repository.pike.location
  repository  = google_artifact_registry_repository.pike.name
  policy_data = data.google_iam_policy.admin.policy_data
}
