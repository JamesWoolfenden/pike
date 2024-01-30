resource "google_api_gateway_api_config" "pike" {
  provider = google-beta
  api      = google_api_gateway_api.pike.api_id
  openapi_documents {
    document {
      path     = "spec.yaml"
      contents = filebase64("spec2.0.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}
