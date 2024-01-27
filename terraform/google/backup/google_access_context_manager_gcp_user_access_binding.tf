resource "google_access_context_manager_gcp_user_access_binding" "gcp_user_access_binding" {
  organization_id = "1231234123"
  group_key       = trimprefix(google_cloud_identity_group.group.id, "groups/")
  access_levels = [
    google_access_context_manager_access_level.access-level.name,
  ]
}
