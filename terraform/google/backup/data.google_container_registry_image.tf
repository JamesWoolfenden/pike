data "google_container_registry_image" "pike" {
  name = "pike"
}

output "image" {
  value = data.google_container_registry_image.pike
}
