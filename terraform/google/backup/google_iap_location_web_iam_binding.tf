resource "google_iap_location_web_iam_binding" "pike" {
  location = "us-central1"

  role = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
