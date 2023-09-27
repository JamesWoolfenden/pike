resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //bigquery
    "bigquery.datasets.create",
    //google_bigtable_instance
    "bigtable.instances.create",
    "bigtable.instances.get",
    "bigtable.clusters.list",
    "bigtable.instances.update",
    "bigtable.instances.delete"

  ]
}
