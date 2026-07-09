data "google_iam_policy" "google_iap_location_web" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_iap_location_web_iam_policy" "pike" {
  location = "us-central1"

  policy_data = data.google_iam_policy.google_iap_location_web.policy_data
}
