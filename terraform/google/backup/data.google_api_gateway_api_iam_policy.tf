data "google_api_gateway_api_iam_policy" "pike" {
  provider = google-beta
  api      = "pike"
}
