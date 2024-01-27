resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "iam.serviceAccounts.create",
    "iam.serviceAccounts.delete",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.update",
    "resourcemanager.projects.get",


    //google_access_context_manager_access_policy
    "accesscontextmanager.accessPolicies.get",
    "accesscontextmanager.accessPolicies.create",
    "accesscontextmanager.accessPolicies.delete",
    "accesscontextmanager.accessPolicies.update",


    //google_access_context_manager_access_levels, google_access_context_manager_access_level
    "accesscontextmanager.accessLevels.get",
    "accesscontextmanager.accessLevels.create",
    "accesscontextmanager.accessLevels.delete",
    "accesscontextmanager.accessLevels.update",


    //google_access_context_manager_authorized_orgs_desc
    "accesscontextmanager.authorizedOrgsDescs.get",
    "accesscontextmanager.authorizedOrgsDescs.create",
    "accesscontextmanager.authorizedOrgsDescs.delete",
    "accesscontextmanager.authorizedOrgsDescs.update",


    //google_access_context_manager_service_perimeter, google_access_context_manager_service_perimeters
    "accesscontextmanager.servicePerimeters.create",
    "accesscontextmanager.servicePerimeters.delete",
    "accesscontextmanager.servicePerimeters.get",
    "accesscontextmanager.servicePerimeters.update",


    //google_access_context_manager_access_policy_iam_binding, google_access_context_manager_access_policy_iam_member,
    //google_access_context_manager_access_policy_iam_policy
    "accesscontextmanager.accessPolicies.getIamPolicy",
    "accesscontextmanager.accessPolicies.setIamPolicy",

    //google_access_context_manager_gcp_user_access_binding
    "accesscontextmanager.gcpUserAccessBindings.create",
    "accesscontextmanager.gcpUserAccessBindings.get",
    "accesscontextmanager.gcpUserAccessBindings.delete",
    "accesscontextmanager.gcpUserAccessBindings.update",


  ]
}
