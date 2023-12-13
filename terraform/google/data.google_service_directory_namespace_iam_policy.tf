data "google_service_directory_namespace_iam_policy" "pike" {
  provider = google-beta
  name     = google_service_directory_namespace.example.name
}

resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "example-namespace"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}
