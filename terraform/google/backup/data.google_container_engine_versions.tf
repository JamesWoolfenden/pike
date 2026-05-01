data "google_container_engine_versions" "pike" {}

output "google_container_engine_versions" {
  value = data.google_container_engine_versions.pike
}
