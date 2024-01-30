resource "google_api_gateway_gateway_iam_policy" "pike" {
  provider    = google-beta
  gateway     = google_api_gateway_gateway.pike.id
  policy_data = data.google_iam_policy.admin.policy_data
}

data "google_iam_policy" "admin" {
  provider = google-beta
  binding {
    role = "roles/apigateway.viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}
