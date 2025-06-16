data "google_beyondcorp_application_iam_policy" "pike" {
  security_gateways_id = "pike"
  application_id       = "pike"
}

output "google_beyondcorp_application_iam_policy" {
  value = data.google_beyondcorp_application_iam_policy.pike
}
