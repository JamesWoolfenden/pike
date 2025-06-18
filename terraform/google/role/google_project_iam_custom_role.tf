resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_cloud_run_v2_worker_pool
    "run.workerpools.create",
    "run.workerpools.get",
    "run.workerpools.update",
    "run.workerpools.delete",


    # //google_cloud_run_v2_service google_cloud_run_service
    # "iam.serviceaccounts.actAs",

    //google_cloud_run_v2_service google_cloud_run_service
    "run.services.create",
    "run.services.get",
    "run.services.delete",
    "run.services.update",

    //google_cloud_run_service_iam_binding
    "run.services.getIamPolicy",
    "run.services.setIamPolicy",

    //google_cloud_run_domain_mapping
    "run.domainmappings.create",
    "run.domainmappings.get",
    "run.domainmappings.delete",
    "run.domainmappings.update",

    //google_cloud_run_v2_service
    "run.operations.get",

    //google_cloud_run_v2_worker_pool_iam_binding
    "run.workerpools.getIamPolicy",
    "run.workerpools.setIamPolicy",

    //google_service_account_iam_member
    "iam.serviceAccounts.getIamPolicy",
    "iam.serviceAccounts.setIamPolicy",

  ]
}
