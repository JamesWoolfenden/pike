resource "google_project_iam_audit_config" "pike" {
  project = "pike-412922"
  service = "allServices"
  audit_log_config {
    log_type = "ADMIN_READ"
  }
  audit_log_config {
    log_type = "DATA_READ"
    # exempted_members = [
    #   "user:james.woolfenden@gmail.com",
    # ]
  }
}
