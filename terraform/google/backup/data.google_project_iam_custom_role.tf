data "google_project_iam_custom_role" "pike" {
  role_id = "terraform_pike"
}

output "google_project_iam_custom_role" {
  value = data.google_project_iam_custom_role.pike
}
