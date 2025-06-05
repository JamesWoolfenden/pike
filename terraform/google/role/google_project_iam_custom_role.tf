resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    # "iam.serviceAccounts.create",
    # "iam.serviceAccounts.delete",
    # "iam.serviceAccounts.get",
    # "iam.serviceAccounts.update",
    "resourcemanager.projects.getIamPolicy",
    "resourcemanager.projects.setIamPolicy",


    # //google_iam_workload_identity_pool
    # "iam.workloadIdentityPools.create",
    # "iam.workloadIdentityPools.delete",
    # "iam.workloadIdentityPools.get",
    # "iam.workloadIdentityPools.update",

    # //google_composer_environment
    # "composer.environments.create",
    # "composer.environments.delete",
    # "composer.environments.get",
    # "composer.environments.update",
    #
    # "iam.serviceAccounts.actAs",
    # "composer.operations.get",
    #
    //google_iam_workload_identity_pool_provider
    "iam.workloadIdentityPoolProviders.create",
    "iam.workloadIdentityPoolProviders.delete",
    "iam.workloadIdentityPoolProviders.get",
    "iam.workloadIdentityPoolProviders.update",
  ]
}
