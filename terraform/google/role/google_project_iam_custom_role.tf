
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [

    "certificatemanager.certs.get",
    "certificatemanager.certs.delete",
    "certificatemanager.dnsauthorizations.get",
    "certificatemanager.dnsauthorizations.delete",
    "certificatemanager.certmapentries.get",
    "certificatemanager.certmapentries.delete",
    "certificatemanager.certmaps.get",
    "certificatemanager.certmaps.delete",
    "certificatemanager.operations.get",

    //google_iam_workload_identity_pool_namespace

    //google_iam_workload_identity_pool_managed_identity
    "iam.workloadIdentityPoolManagedIdentities.create",
    "iam.workloadIdentityPoolManagedIdentities.delete",
    "iam.workloadIdentityPoolManagedIdentities.get",
    "iam.workloadIdentityPoolManagedIdentities.update",

    //google_iam_principal_access_boundary_policy
    # "iam.principalaccessboundarypolicies.create",
    # "iam.principalaccessboundarypolicies.get",
    # "iam.principalaccessboundarypolicies.update",
    # "iam.principalaccessboundarypolicies.delete",

    //google_iam_oauth_client
    "iam.oauthClients.create",
    "iam.oauthClients.delete",
    "iam.oauthClients.update",
    "iam.oauthClients.get",

    # "iam.googleapis.com/workforcePoolProviders.create",
    # "iam.googleapis.com/workforcePoolProviders.delete",
    # "iam.googleapis.com/workforcePoolProviders.get",
    # "iam.googleapis.com/workforcePoolProviders.update",
    # "iam.googleapis.com/workforcePools.create",
    # "iam.googleapis.com/workforcePools.delete",
    # "iam.googleapis.com/workforcePools.get",
    # "iam.googleapis.com/workforcePools.update",
    "iam.serviceAccounts.create",
    "iam.serviceAccounts.delete",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.update",
    "iam.workloadIdentityPools.create",
    "iam.workloadIdentityPools.delete",
    "iam.workloadIdentityPools.get",
    "iam.workloadIdentityPools.update",
    # "resourcemanager.organizations.get",
    # "resourcemanager.projects.create",
    "resourcemanager.projects.delete",
    "resourcemanager.projects.get",
    "resourcemanager.projects.update",
    "storage.buckets.get"
  ]
}
