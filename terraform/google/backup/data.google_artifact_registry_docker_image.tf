data "google_artifact_registry_docker_image" "pike" {
  location      = "us-central1"
  repository_id = "pike"
  image_name    = "pike"
}

output "google_artifact_registry_docker_image" {
  value = data.google_artifact_registry_docker_image.pike
}
