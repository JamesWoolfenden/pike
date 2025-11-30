resource "google_project_iam_member_remove" "pike" {
  project = "pike-477416"
  role    = "roles/compute.networkUser"
  member  = "user:crwwoolfenden@gmail.com"
}
