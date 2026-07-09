# resource "google_iap_agent_registry_iam_binding" "pike" {
#   location = "us-central1"
#
#   role = "roles/viewer"
#   members = [
#     "user:james.woolfenden@gmail.com",
#   ]
# }


resource "google_iap_agent_registry_iam_binding" "binding" {
  # project = google_project_service.project_service.project
  location = "us-central1"
  role     = "roles/iap.egressor"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
