data "google_privateca_certificate_template_iam_policy" "pike" {
  certificate_template = google_privateca_certificate_template.default.id
}

output "google_privateca_certificate_template_iam_policy" {
  value = data.google_privateca_certificate_template_iam_policy.pike
}
