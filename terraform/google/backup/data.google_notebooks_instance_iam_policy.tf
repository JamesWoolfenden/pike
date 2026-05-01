data "google_notebooks_instance_iam_policy" "pike" {
  instance_name = "pike"
  location      = "us-central1"
}

output "google_notebooks_instance_iam_policy" {
  value = data.google_notebooks_instance_iam_policy.pike
}
