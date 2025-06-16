data "google_memorystore_instance" "pike" {
  instance_id = "pike"
}

output "google_memorystore_instance" {
  value = data.google_memorystore_instance.pike
}
