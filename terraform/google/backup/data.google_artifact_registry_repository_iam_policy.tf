data "google_artifact_registry_repository_iam_policy" "pike" {
  project    = data.google_artifact_registry_repository.pike.project
  location   = data.google_artifact_registry_repository.pike.location
  repository = data.google_artifact_registry_repository.pike.name
}

output "policy" {
  value = data.google_artifact_registry_repository_iam_policy.pike
}
