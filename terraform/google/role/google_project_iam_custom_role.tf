resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [

    //google_project_service
    "serviceusage.services.get",
    "serviceusage.services.list",
    "serviceusage.services.enable",
    "serviceusage.services.disable",

    "resourcemanager.projects.get",
    #    "iam.serviceAccounts.list",
    #    "iam.serviceAccounts.setIamPolicy",
    #    "iam.serviceAccounts.getIamPolicy",
    #    "iam.serviceAccounts.undelete"



  ]
}
