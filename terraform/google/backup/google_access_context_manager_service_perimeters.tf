resource "google_access_context_manager_service_perimeters" "pike" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"

  service_perimeters {
    name  = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/"
    title = ""
    status {
      restricted_services = ["storage.googleapis.com"]
    }
  }

  service_perimeters {
    name  = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/"
    title = ""
    status {
      restricted_services = ["bigtable.googleapis.com"]
    }
  }
}


resource "google_access_context_manager_access_policy" "access-policy" {
  parent = data.google_organization.this.id
  title  = "pike"
}

data "google_organization" "this" {
  organization = "12345354"
}
