data "google_cloud_run_v2_job_iam_policy" "pike" {
  name = "pike"
}

output "policy2" {
  value = data.google_cloud_run_v2_job_iam_policy.pike
}
