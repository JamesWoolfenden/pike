resource "google_project_iam_member_remove" "pike" {
  project = "pike-412922"
  role    = "roles/compute.networkUser"
  member  = "user:crwwoolfenden@gmail.com"
}
