resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "resourcemanager.projects.get",
    //google_cloud_run_v2_job
    "run.jobs.create",
    "run.jobs.get",
    "run.operations.get",
    "run.jobs.delete",
    "run.jobs.update",
    //google_cloud_scheduler_job
    "cloudscheduler.jobs.create",
    "cloudscheduler.jobs.enable",
    "cloudscheduler.jobs.get",
    "cloudscheduler.jobs.delete",
    "cloudscheduler.jobs.update"
  ]
}
