resource "google_compute_project_metadata" "pike" {
  metadata = {
    pike-test-key = "pike-test-value"
  }
}
