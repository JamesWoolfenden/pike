data "google_secure_source_manager_repository_iam_policy" "pike" {
}

output "google_secure_source_manager_repository_iam_policy" {
  value = data.google_secure_source_manager_repository_iam_policy.pike
}
