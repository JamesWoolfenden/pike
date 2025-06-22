data "google_iam_policy" "admin2" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_data_catalog_tag_template_iam_policy" "policy" {
  tag_template = google_data_catalog_tag_template.pike.name
  policy_data  = data.google_iam_policy.admin2.policy_data
}
