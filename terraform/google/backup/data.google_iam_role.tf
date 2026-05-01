data "google_iam_role" "roleinfo" {
  name = "roles/compute.viewer"
}

output "google_iam_role" {
  value = data.google_iam_role.roleinfo.included_permissions
}
