data "google_container_attached_install_manifest" "pike" {
  cluster_id       = "pike"
  location         = "us-central1"
  platform_version = "pike"
  project          = "pike-gcp"
}
