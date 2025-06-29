resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //data.google_project_ancestry google_project
    "resourcemanager.projects.get",

    //google_project
    "resourcemanager.organizations.get",
    "resourcemanager.projects.create",
    "resourcemanager.projects.delete",
    "resourcemanager.projects.get",
    "resourcemanager.projects.update",

    //google_project_access_approval_settings
    "accessapproval.settings.delete",
    "accessapproval.settings.get",
    "accessapproval.settings.create",
    "accessapproval.settings.update",

    # //google_project_organization_policy
    "orgpolicy.policy.set",
    "orgpolicy.policy.get",

    //google_project_iam_member_remove
    "resourcemanager.projects.setIamPolicy",
    "resourcemanager.projects.getIamPolicy",

  ]
}
