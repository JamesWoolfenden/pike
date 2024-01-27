resource "google_access_context_manager_access_policy_iam_member" "pike" {
  name   = google_access_context_manager_access_policy.access-policy.name
  role   = "roles/accesscontextmanager.policyAdmin"
  member = "user:crwoolfenden@gmail.com"
}
