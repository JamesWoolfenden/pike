resource "google_project_iam_custom_role" "pike" {
  project     = "pike-361314"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "resourcemanager.projects.setIamPolicy",
    "resourcemanager.projects.getIamPolicy",
    #        "iam.serviceAccounts.get",
    #    "iam.serviceAccounts.setIamPolicy",
    #    "iam.serviceAccounts.getIamPolicy"
  ]
}
