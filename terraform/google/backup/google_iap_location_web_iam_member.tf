resource "google_iap_location_web_iam_member" "pike" {
  location = "us-central1"

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}
