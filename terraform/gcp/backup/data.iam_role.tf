data "google_iam_role" "roleinfo" {
  name = "roles/compute.viewer"
}

output "the_role_permissions" {
  value = data.google_iam_role.roleinfo.included_permissions
}
