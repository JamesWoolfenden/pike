data "google_project_iam_policy" "pike" {
}

output "google_project_iam_policy" {
  value = data.google_project_iam_policy.pike
}
