data "google_iam_workload_identity_pool_iam_policy" "pike" {
  provider                  = google-beta
  workload_identity_pool_id = "gitlab"
}

output "google_iam_workload_identity_pool_iam_policy" {
  value = data.google_iam_workload_identity_pool_iam_policy.pike
}
