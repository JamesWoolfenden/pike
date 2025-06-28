data "google_scc_source_iam_policy" "pike" {
}

output "google_scc_source_iam_policy" {
  value = data.google_scc_source_iam_policy.pike
}
