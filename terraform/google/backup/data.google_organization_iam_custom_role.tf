data "google_organization_iam_custom_role" "pike" {
  role_id = "pike"
  org_id  = "123456789"
}


output "google_organization_iam_custom_role" {
  value = data.google_organization_iam_custom_role.pike
}
