resource "google_project_iam_custom_role" "pike" {
  project     = "pike-361314"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "iam.serviceAccounts.actAs",
    "cloudfunctions.functions.create",
    "cloudfunctions.operations.get",
    "cloudfunctions.functions.get",
    "cloudfunctions.functions.delete",
    "cloudfunctions.functions.update",

    "cloudfunctions.functions.setIamPolicy",
    "cloudfunctions.functions.getIamPolicy"
  ]
}
