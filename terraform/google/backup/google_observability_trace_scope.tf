resource "google_observability_trace_scope" "pike" {
  provider       = google-beta
  trace_scope_id = "test-scope"
  location       = "global"
  resource_names = [
  "projects/123456"]
  description = "A trace scope configured with Terraform"
}
