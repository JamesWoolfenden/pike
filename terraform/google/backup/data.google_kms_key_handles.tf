data "google_kms_key_handles" "pike" {
  provider               = google-beta
  location               = "us-central1"
  resource_type_selector = "storage.googleapis.com/Bucket"
}

output "google_kms_key_handles" {
  value = data.google_kms_key_handles.pike
}
