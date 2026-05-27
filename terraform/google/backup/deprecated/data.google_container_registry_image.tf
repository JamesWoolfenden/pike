data "google_container_registry_image" "pike" {
  name = "pike"
}

output "google_container_registry_image" {
  value = data.google_container_registry_image.pike
}
