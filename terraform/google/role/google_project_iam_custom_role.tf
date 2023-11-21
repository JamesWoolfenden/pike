resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_cloud_run_v2_service_iam_policy, google_cloud_run_v2_job_iam_policy, google_cloud_run_service_iam_policy
    "run.services.getIamPolicy",
    //google_cloud_run_service
    "run.services.get",
    //google_cloud_run_locations
    "run.locations.list",
    //google_cloud_run_v2_job
    "run.jobs.get",
    //google_cloud_run_v2_job_iam_policy
    "run.jobs.getIamPolicy"
  ]
}
