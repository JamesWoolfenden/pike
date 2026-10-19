resource "google_service_directory_namespace" "pike" {
  namespace_id = "pike-test-namespace"
  location     = "europe-west2"
  labels = {
    test = "modify"
  }
}
