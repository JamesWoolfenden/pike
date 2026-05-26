data "google_notebooks_runtime_iam_policy" "pike" {
  runtime_name = "pike"
  location     = "us-central1"
}

output "google_notebooks_runtime_iam_policy" {
  value = data.google_notebooks_runtime_iam_policy.pike
}
