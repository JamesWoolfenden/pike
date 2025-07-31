data "google_project_organization_policy" "pike" {
  project    = "pike-gcp"
  constraint = "constraints/serviceuser.services"
}

output "version" {
  value = data.google_project_organization_policy.pike.version
}
