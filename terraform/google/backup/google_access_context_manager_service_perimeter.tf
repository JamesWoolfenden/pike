resource "google_access_context_manager_service_perimeter" "service-perimeter-resource" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/restrict_all"
  title  = "restrict_all"
  status {
    restricted_services = ["storage.googleapis.com"]
  }

  lifecycle {
    ignore_changes = [status[0].resources]
  }
}
