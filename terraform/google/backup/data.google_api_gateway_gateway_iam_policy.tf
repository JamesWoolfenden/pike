data "google_api_gateway_gateway_iam_policy" "pike" {
  provider = google-beta
  gateway  = "pike"
}
