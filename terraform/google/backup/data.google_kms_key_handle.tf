data "google_kms_key_handle" "pike" {
  name     = "pike"
  location = "us-central1"
  provider = google-beta
}

output "google_kms_key_handle" {
  value = data.google_kms_key_handle.pike
}
