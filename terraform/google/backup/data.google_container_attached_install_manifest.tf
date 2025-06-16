data "google_container_attached_install_manifest" "pike" {
  cluster_id       = "pike"
  project          = "pike-412922"
  platform_version = "1234"
  location         = "us-central1"
}

output "google_container_attached_install_manifest" {
  value = data.google_container_attached_install_manifest.pike
}
