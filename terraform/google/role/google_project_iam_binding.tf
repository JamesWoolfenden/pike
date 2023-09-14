resource "google_project_iam_binding" "pike" {
  project = "pike-gcp"
  role    = google_project_iam_custom_role.pike.id

  members = [
    "serviceAccount:${google_service_account.pike.email}",
  ]
}
