data "google_scc_v2_organization_source_iam_policy" "pike" {
  source       = google_scc_v2_organization_source.custom_source.name
  organization = "pike"
}

output "google_scc_v2_organization_source_iam_policy" {
  value = data.google_scc_v2_organization_source_iam_policy.pike
}
