data "google_project_iam_policy" "pike" {
  project = "pike-gcp"
}

output "policy" {
  value = data.google_project_iam_policy.pike
}
