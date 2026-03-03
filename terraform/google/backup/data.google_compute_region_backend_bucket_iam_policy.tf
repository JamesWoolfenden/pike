data "google_compute_region_backend_bucket_iam_policy" "pike" {
  provider = google-beta
  name     = "terraform-pike-bucket-tfstate"
}

output "google_compute_region_backend_bucket_iam_policy" {
  value = data.google_compute_region_backend_bucket_iam_policy.pike
}
