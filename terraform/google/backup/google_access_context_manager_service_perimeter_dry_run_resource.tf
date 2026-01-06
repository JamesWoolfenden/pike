resource "google_access_context_manager_service_perimeter_dry_run_resource" "service-perimeter-dry-run-resource" {
  perimeter_name = google_access_context_manager_service_perimeter.service-perimeter-dry-run-resource.name
  resource       = "projects/987654321"
}

resource "google_access_context_manager_service_perimeter" "service-perimeter-dry-run-resource" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/restrict_all"
  title  = "restrict_all"
  spec {
    restricted_services = ["storage.googleapis.com"]
  }
  use_explicit_dry_run_spec = true
  lifecycle {
    ignore_changes = [spec[0].resources]
  }
}
