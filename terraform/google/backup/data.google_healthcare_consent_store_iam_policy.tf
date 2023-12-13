data "google_healthcare_consent_store_iam_policy" "pike" {
  consent_store_id = "pike"
  dataset          = "pike"
}

output "policy5" {
  value = data.google_healthcare_consent_store_iam_policy.pike
}
