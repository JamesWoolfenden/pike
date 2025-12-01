data "google_container_azure_versions" "pike" {
  location = "us-west1"
  project  = "pike-477416"
}

output "google_container_azure_versions" {
  value = data.google_container_azure_versions.pike
}
