data "google_kms_autokey_config" "pike" {
  provider = google-beta
  folder   = "pike"
}

output "google_kms_autokey_config" {
  value = data.google_kms_autokey_config.pike
}
