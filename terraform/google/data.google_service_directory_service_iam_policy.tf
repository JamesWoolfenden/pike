data "google_service_directory_service_iam_policy" "pike" {
  provider = google-beta
  name     = google_service_directory_namespace.example.name
}

resource "google_service_directory_service" "example" {
  provider   = google-beta
  service_id = "example-service"
  namespace  = google_service_directory_namespace.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}
