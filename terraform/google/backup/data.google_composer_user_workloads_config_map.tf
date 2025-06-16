data "google_composer_user_workloads_config_map" "pike" {
  name        = "pike"
  environment = "pike"
}

output "google_composer_user_workloads_config_map" {
  value = data.google_composer_user_workloads_config_map.pike
}
