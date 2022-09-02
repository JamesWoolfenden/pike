resource "google_project_iam_custom_role" "pike" {
  project     = "examplea"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "iam.roles.get",
    "iam.roles.create",

    "iam.roles.update",

    "iam.roles.delete"
  ]
}
