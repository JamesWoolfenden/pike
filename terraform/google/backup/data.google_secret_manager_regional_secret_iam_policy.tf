data "google_secret_manager_regional_secret_iam_policy" "pike" {
  secret_id = "secretname"
}

output "google_secret_manager_regional_secret_iam_policy" {
  value = data.google_secret_manager_regional_secret_iam_policy.pike
}
