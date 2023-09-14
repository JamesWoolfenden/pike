data "google_kms_key_ring" "pike" {
  name     = "pike"
  location = "europe-west1"
}

output "ring" {
  value = data.google_kms_key_ring.pike
}
