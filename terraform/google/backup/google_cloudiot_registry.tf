resource "google_cloudiot_registry" "pike" {
  provider = google-beta
  name     = "cloudiot-registry"
}
