data "google_api_gateway_api_config_iam_policy" "pike" {
  provider   = google-beta
  api_config = "pike"
  api        = "pike"
}
