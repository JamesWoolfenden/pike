resource "google_iam_access_boundary_policy" "pike" {
  parent       = urlencode("cloudresourcemanager.googleapis.com/projects/${google_project.project.project_id}")
  name         = "my-ab-policy"
  display_name = "My AB policy"
  rules {
    description = "AB rule"
    access_boundary_rule {
      available_resource    = "*"
      available_permissions = ["*"]
      availability_condition {
        title      = "Access level expr"
        expression = "request.matchAccessLevels('${google_project.project.org_id}', ['${google_access_context_manager_access_level.test-access.name}'])"
      }
    }
  }
}
