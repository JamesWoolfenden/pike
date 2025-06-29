data "google_runtimeconfig_config" "pike" {
}

output "google_runtimeconfig_config" {
  value = data.google_runtimeconfig_config.pike
}
