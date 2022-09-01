resource "google_project_iam_binding" "pike" {
  project = "examplea"
  role    = google_project_iam_custom_role.pike.id

  members = [
    "serviceAccount:${google_service_account.pike.email}",
  ]
}
