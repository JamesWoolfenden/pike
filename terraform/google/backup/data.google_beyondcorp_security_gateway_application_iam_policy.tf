data "google_beyondcorp_security_gateway_application_iam_policy" "pike" {
  application_id      = "pike"
  security_gateway_id = "pike"
}

output "google_beyondcorp_security_gateway_application_iam_policy" {
  value = data.google_beyondcorp_security_gateway_application_iam_policy.pike
}
