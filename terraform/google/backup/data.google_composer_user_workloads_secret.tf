data "google_composer_user_workloads_secret" "pike" {
  environment = "pike"
  name        = "pike"
}

output "google_composer_user_workloads_secret" {
  value = data.google_composer_user_workloads_secret.pike
}
