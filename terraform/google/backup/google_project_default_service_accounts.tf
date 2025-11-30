resource "google_project_default_service_accounts" "my_project" {
  project        = "pike-477416"
  action         = "DISABLE"
  restore_policy = "REVERT"
}
