resource "google_project_access_approval_settings" "pike" {
  project_id          = "pike-plus"
  notification_emails = ["james.woolfenden@gmail.com"]

  enrolled_services {
    cloud_product    = "all"
    enrollment_level = "BLOCK_ALL"
  }
}
