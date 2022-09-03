resource "google_kms_key_ring" "pike" {
  name     = "pike"
  location = "europe-west1"
}

# key rings cannot be deleted
# terraform import google_kms_key_ring.pike projects/pike-361314/locations/europe-west1/keyRings/pike
