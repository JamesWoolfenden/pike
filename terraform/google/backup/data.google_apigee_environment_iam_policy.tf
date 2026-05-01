data "google_apigee_environment_iam_policy" "pike" {
  env_id = "123456"
  org_id = "organization/1231234"
}

output "google_apigee_environment_iam_policy" {
  value = data.google_apigee_environment_iam_policy.pike
}
