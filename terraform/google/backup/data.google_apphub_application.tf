data "google_apphub_application" "pike" {
  location       = "us-central1"
  application_id = "pike"
  project        = "pike-477416"
}

output "google_apphub_application" {
  value = data.google_apphub_application.pike
}
