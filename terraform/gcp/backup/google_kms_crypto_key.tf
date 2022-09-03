resource "google_kms_crypto_key" "pike" {
  name            = "pike-key-keep"
  key_ring        = "projects/pike-361314/locations/europe-west1/keyRings/pike"
  rotation_period = "100000s"

  purpose = "ENCRYPT_DECRYPT"
  labels = {
    "pike" = "permissions"
    update = "possible"
  }
  # destroy_scheduled_duration = "86400s"

  # lifecycle {
  #   prevent_destroy = true
  # }
}
