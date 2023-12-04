data "google_notebooks_runtime_iam_policy" "pike" {
  runtime_name = "pike"
  location     = "us-central1"
}

output "policy3" {
  value = data.google_notebooks_runtime_iam_policy.pike
}
