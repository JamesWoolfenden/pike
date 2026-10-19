resource "google_kms_project_autokey_config" "pike" {
  key_project_resolution_mode = "RESOURCE_PROJECT"
}