data "google_dataproc_job_iam_policy" "pike" {
  job_id = "pike"
}

output "policy6" {
  value = data.google_dataproc_job_iam_policy.pike
}
