resource "google_api_gateway_api_config_iam_policy" "pike" {
  provider    = google-beta
  policy_data = data.google_iam_policy.admin.policy_data
  api         = google_api_gateway_api_config.pike.api
  api_config  = google_api_gateway_api_config.pike.api_config_id
}
