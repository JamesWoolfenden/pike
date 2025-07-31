data "google_secure_source_manager_instance_iam_policy" "pike" {
  instance_id = "pike"
}

output "google_secure_source_manager_instance_iam_policy" {
  value = data.google_secure_source_manager_instance_iam_policy.pike
}
