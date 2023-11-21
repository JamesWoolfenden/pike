data "google_compute_disk" "pike" {
  name    = "pike"
  project = "pike-gcp"
  zone    = "europe-west1-b"
}
