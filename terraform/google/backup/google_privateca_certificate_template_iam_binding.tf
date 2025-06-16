resource "google_privateca_certificate_template_iam_binding" "binding" {
  certificate_template = google_privateca_certificate_template.default.id
  role                 = "roles/privateca.templateUser"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
