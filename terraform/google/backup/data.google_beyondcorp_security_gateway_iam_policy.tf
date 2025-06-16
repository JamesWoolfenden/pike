data "google_beyondcorp_security_gateway_iam_policy" "pike" {
  security_gateway_id = "pike"
  location            = "us-central1"
}

output "google_beyondcorp_security_gateway_iam_policy" {
  value = data.google_beyondcorp_security_gateway_iam_policy.pike
}
