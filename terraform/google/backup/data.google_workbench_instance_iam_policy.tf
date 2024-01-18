data "google_workbench_instance_iam_policy" "pike" {
  project  = "pike-gcp"
  location = "us-central1"
  name     = "pike"
}
