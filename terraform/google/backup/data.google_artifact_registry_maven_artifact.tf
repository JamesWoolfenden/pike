data "google_artifact_registry_maven_artifact" "pike" {
  artifact_id   = "pike"
  group_id      = "pike"
  location      = "pike"
  repository_id = "pike"
}

output "google_artifact_registry_maven_artifact" {
  value = data.google_artifact_registry_maven_artifact.pike
}
