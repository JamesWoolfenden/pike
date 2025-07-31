data "google_runtimeconfig_config_iam_policy" "pike" {
  provider = google-beta
  config   = "pike"
}
