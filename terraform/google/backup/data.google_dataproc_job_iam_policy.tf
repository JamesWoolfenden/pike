data "google_dataproc_job_iam_policy" "pike" {
  job_id = "pike"
}

output "google_dataproc_job_iam_policy" {
  value = data.google_dataproc_job_iam_policy.pike
}
