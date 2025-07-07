
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "secretmanager.secrets.create",
    "secretmanager.secrets.delete",
    "secretmanager.secrets.get",
    "secretmanager.secrets.update",
    "secretmanager.versions.access",
    "secretmanager.versions.add",
    "secretmanager.versions.destroy",
    "secretmanager.versions.enable",
    "secretmanager.versions.get",
    "source.repos.create",
    "source.repos.delete",
    "source.repos.get",
    "source.repos.updateRepoConfig",
    "storage.buckets.get",
    "iam.serviceAccounts.create",
    "iam.serviceAccounts.delete",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.list",
    "iam.serviceAccounts.update",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.list",
    "storage.objects.create",
    "storage.objects.delete",
    "pubsub.topics.create",
    "pubsub.topics.delete",
    "pubsub.topics.get",
    "pubsub.topics.update",


    "dataform.repositories.create",
    "dataform.repositories.get",
    "dataform.repositories.delete",

    "dataform.releaseConfigs.create",
    "dataform.releaseConfigs.get",
    "dataform.releaseConfigs.update",
    "dataform.releaseConfigs.delete",

    "dataform.workflowConfigs.create",
    "dataform.workflowConfigs.get",
    "dataform.workflowConfigs.update",
    "dataform.workflowConfigs.delete",


    //dataflow_job
    "resourcemanager.projects.get",
    "dataflow.jobs.create",
    "dataflow.jobs.get"

  ]
}
