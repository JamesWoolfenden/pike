resource "google_api_gateway_api_iam_policy" "pike" {
  provider    = google-beta
  policy_data = data.google_iam_policy.admin2.policy_data
  api         = google_api_gateway_api.pike.api_id
}

data "google_iam_policy" "admin2" {
  provider = google-beta
  binding {
    role = "roles/apigateway.viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}
