data "google_artifact_registry_docker_images" "pike" {
}

output "google_artifact_registry_docker_images" {
  value = data.google_artifact_registry_docker_images.pike
}
