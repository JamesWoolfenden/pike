resource "google_project_iam_custom_role" "pike" {
  project     = "pike-361314"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "container.clusters.create",
    "container.operations.get",
    "container.clusters.get",
    "compute.instanceGroupManagers.get",
    "container.clusters.delete",

    "iam.serviceAccounts.actAs",
    "container.clusters.update"
  ]
}
