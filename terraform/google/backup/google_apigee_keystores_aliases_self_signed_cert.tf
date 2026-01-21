resource "google_apigee_keystores_aliases_self_signed_cert" "pike" {
  alias       = "alias"
  org_id      = google_apigee_organization.org.id
  environment = "pike"
  keystore    = "pike-12345"
  sig_alg     = "SHA512withRSA"
  subject {
    common_name  = "selfsigned_example"
    country_code = "US"
    locality     = "TX"
    org          = "CCE"
    org_unit     = "PSO"
  }
}
