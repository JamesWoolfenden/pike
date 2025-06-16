data "google_kms_key_rings" "pike" {
  location = "us-central1"
}

output "google_kms_key_rings" {
  value = data.google_kms_key_rings.pike
}
