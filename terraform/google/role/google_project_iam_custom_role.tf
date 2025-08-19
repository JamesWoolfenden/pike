
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "managedkafka.clusters.create",
    "managedkafka.clusters.delete",
    "managedkafka.clusters.get",
    "managedkafka.clusters.update",
    "managedkafka.operations.get",
    "resourcemanager.projects.get",
    "storage.buckets.get",

    "managedkafka.acls.create",
    "managedkafka.acls.get",
    "managedkafka.acls.update",
    "managedkafka.acls.delete",

  ]
}
