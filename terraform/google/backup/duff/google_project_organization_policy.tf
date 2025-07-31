resource "google_project_organization_policy" "pike" {
  resource "google_project_iam_policy" "project" {
    project     = "your-project-id"
    policy_data = data.google_iam_policy.admin.policy_data
  }

  data "google_iam_policy" "admin" {
    binding {
      role = "roles/editor"

      members = [
        "user:james.woolfenden@gmail.com",
      ]
    }
  }
}
