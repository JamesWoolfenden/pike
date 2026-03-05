data "google_artifact_registry_tag" "pike" {
  package_name  = "pike"
  tag_name      = "pike"
  repository_id = "pike"
  location      = "us-central1"
}

output "google_artifact_registry_tag" {
  value = data.google_artifact_registry_tag.pike
}
