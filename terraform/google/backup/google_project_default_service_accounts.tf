resource "google_project_default_service_accounts" "my_project" {
  project        = "pike-412922"
  action         = "DISABLE"
  restore_policy = "REVERT"
}
