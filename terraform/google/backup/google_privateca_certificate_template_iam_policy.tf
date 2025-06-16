data "google_iam_policy" "admin" {
  binding {
    role = "roles/privateca.templateUser"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_privateca_certificate_template_iam_policy" "policy" {
  certificate_template = google_privateca_certificate_template.default.id
  policy_data          = data.google_iam_policy.admin.policy_data
}
