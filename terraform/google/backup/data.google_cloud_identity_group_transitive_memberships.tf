data "google_cloud_identity_group_transitive_memberships" "pike" {
  group = "pike"
}

output "google_cloud_identity_group_transitive_memberships" {
  value = data.google_cloud_identity_group_transitive_memberships.pike
}
