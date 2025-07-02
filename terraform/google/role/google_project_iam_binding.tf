resource "google_project_iam_binding" "pike" {
  project = "pike-412922"
  role    = google_project_iam_custom_role.terraform_pike.id

  members = [
    "serviceAccount:${google_service_account.pike.email}",
  ]
}
