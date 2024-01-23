resource "google_service_directory_service" "pike" {
  provider   = google-beta
  service_id = "example-service"
  namespace  = google_service_directory_namespace.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}
