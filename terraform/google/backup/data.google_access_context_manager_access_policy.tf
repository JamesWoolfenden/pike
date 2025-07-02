data "google_access_context_manager_access_policy" "pike" {
}

output "google_access_context_manager_access_policy" {
  value = data.google_access_context_manager_access_policy.pike
}
