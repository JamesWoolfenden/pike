resource "google_project_iam_binding" "pike" {
  project = "pike-361314"
  role    = "roles/editor"

  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
