data "google_project_organization_policy" "pike" {
}

output "google_project_organization_policy" {
  value = data.google_project_organization_policy.pike
}
