data "google_iam_testable_permissions" "pike" {
  full_resource_name = "//cloudresourcemanager.googleapis.com/projects/pike"
}

output "google_iam_testable_permissions" {
  value = data.google_iam_testable_permissions.pike
}
