data "google_cloud_run_v2_job" "pike" {
  provider = google-beta
  name     = "pike"
  project  = "pike"
  location = "europe-west2"
}

output "job" {
  value = data.google_cloud_run_v2_job.pike
}
