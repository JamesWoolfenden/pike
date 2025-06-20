resource "google_kms_key_ring_import_job" "pike" {
  key_ring      = google_kms_key_ring.keyring.id
  import_job_id = "my-import-job2"

  import_method    = "RSA_OAEP_3072_SHA1_AES_256"
  protection_level = "SOFTWARE"
}
