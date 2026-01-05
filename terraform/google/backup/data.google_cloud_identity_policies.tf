data "google_cloud_identity_policies" "pike" {
}

output "google_cloud_identity_policies" {
  value = data.google_cloud_identity_policies.pike
}
