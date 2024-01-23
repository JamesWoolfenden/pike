resource "google_service_directory_endpoint" "pike" {
  provider    = google-beta
  endpoint_id = "example-endpoint"
  service     = google_service_directory_service.pike.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }

  address = "1.2.3.4"
  port    = 5353
}
