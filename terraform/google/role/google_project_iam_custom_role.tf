
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [

    //google_cloud_tasks_queue


    //google_clouddeploy_custom_target_type


    //google_clouddeploy_delivery_pipeline

    //google_clouddeploy_deploy_policy

    //google_clouddeploy_target



    //google_clouddeploy_automation


    "iam.serviceAccounts.actAs",

  ]
}
