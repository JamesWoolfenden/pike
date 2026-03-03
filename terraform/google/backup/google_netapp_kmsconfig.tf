resource "google_kms_key_ring" "pike_keyring" {
  name     = "pike-netapp-keyring"
  location = "europe-west2"
}

resource "google_kms_crypto_key" "pike_key" {
  name     = "pike-netapp-key"
  key_ring = google_kms_key_ring.pike_keyring.id
}

resource "google_netapp_kmsconfig" "pike" {
  name            = "pike-kmsconfig"
  location        = "europe-west2"
  crypto_key_name = google_kms_crypto_key.pike_key.id
  description     = "Pike test KMS config - updated"
}
