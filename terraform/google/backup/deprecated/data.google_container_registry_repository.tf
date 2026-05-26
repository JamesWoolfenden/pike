data "google_container_registry_repository" "pike" {}

output "google_container_registry_repository" {
  value = data.google_container_registry_repository.pike
}
