resource "google_privateca_certificate_template_iam_member" "member" {
  certificate_template = google_privateca_certificate_template.default.id
  role                 = "roles/privateca.templateUser"
  member               = "user:james.woolfenden@gmail.com"
}
