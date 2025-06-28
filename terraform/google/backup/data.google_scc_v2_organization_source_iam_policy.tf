data "google_scc_v2_organization_source_iam_policy" "pike" {
}

output "google_scc_v2_organization_source_iam_policy" {
  value = data.google_scc_v2_organization_source_iam_policy.pike
}
