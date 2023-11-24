data "google_container_engine_versions" "pike" {}

output "versions" {
  value = data.google_container_engine_versions.pike
}
