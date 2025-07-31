resource "google_project_iam_member_remove" "foo" {
  role    = "roles/editor"
  project = "pike-412922"
  member  = "serviceAccount:dumbass@developer.gserviceaccount.com"
}
