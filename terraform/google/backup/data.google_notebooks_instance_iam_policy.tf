data "google_notebooks_instance_iam_policy" "pike" {
  instance_name = "pike"
  location      = "us-central1"
}

output "policy2" {
  value = data.google_notebooks_instance_iam_policy.pike
}
