data "google_runtimeconfig_variable" "pike" {
}

output "google_runtimeconfig_variable" {
  value = data.google_runtimeconfig_variable.pike
}
