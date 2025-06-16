resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    # google_cloud_identity_group_transitive_memberships.json
    # google_firebase_web_app_config.json
  ]
}
