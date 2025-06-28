data "google_runtimeconfig_config_iam_policy" "pike" {
}

output "google_runtimeconfig_config_iam_policy" {
  value = data.google_runtimeconfig_config_iam_policy.pike
}
