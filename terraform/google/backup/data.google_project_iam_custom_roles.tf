data "google_project_iam_custom_roles" "pike" {

}

output "google_project_iam_custom_roles" {
  value = data.google_project_iam_custom_roles.pike
}
