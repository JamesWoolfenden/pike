data "google_cloud_identity_group_transitive_memberships" "pike" {
  group = "groups/123eab45c6defghi"
}

output "google_cloud_identity_group_transitive_memberships" {
  value = data.google_cloud_identity_group_transitive_memberships.pike
}
