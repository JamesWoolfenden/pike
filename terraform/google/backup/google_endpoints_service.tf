resource "google_endpoints_service" "pike" {
  service_name   = "api-name.endpoints.project-id.cloud.goog"
  project        = "project-id"
  openapi_config = file("openapi_spec.yml")
}
