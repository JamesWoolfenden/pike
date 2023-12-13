data "google_endpoints_service_consumers_iam_policy" "pike" {
  consumer_project = "pike"
  service_name     = "pike"
}
