data "google_organization_iam_custom_roles" "pike" {
  org_id       = "1234567890"
  show_deleted = true
  view         = "FULL"
}

output "google_organization_iam_custom_roles" {
  value = data.google_organization_iam_custom_roles.pike
}
