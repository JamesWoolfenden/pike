data "google_iam_policy" "admin" {
  binding {
    role = "roles/accesscontextmanager.policyAdmin"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_access_context_manager_access_policy_iam_policy" "policy" {
  name        = google_access_context_manager_access_policy.access-policy.name
  policy_data = data.google_iam_policy.admin.policy_data
}
