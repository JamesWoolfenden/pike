data "google_container_registry_repository" "pike" {}

output "repo" {
  value = data.google_container_registry_repository.pike
}
