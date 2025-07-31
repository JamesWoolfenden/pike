data "google_runtimeconfig_variable" "pike" {
  provider = google-beta
  name     = "pike"
  parent   = "organization/1223435"
}
