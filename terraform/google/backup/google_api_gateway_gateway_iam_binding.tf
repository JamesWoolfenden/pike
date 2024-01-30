resource "google_api_gateway_gateway_iam_binding" "pike" {
  provider = google-beta
  gateway  = google_api_gateway_gateway.pike.id
  role     = "roles/apigateway.viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
