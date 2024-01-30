resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_spanner_instance
    "spanner.instances.create",
    "spanner.instances.delete",
    "spanner.instances.update",
    "spanner.instances.get",
    "spanner.instanceOperations.get",

    //google_spanner_instance_iam_policy,google_spanner_instance_iam_member,google_spanner_instance_iam_binding
    "spanner.instances.setIamPolicy",
    "spanner.instances.getIamPolicy",

    //google_spanner_database
    "spanner.databases.create",
    "spanner.databases.drop",
    "spanner.databases.updateDdl",
    "spanner.databases.update",
    "spanner.databases.get",
    "spanner.databaseOperations.get",

    //google_spanner_database_iam_policy,google_spanner_database_iam_member,google_spanner_database_iam_binding
    "spanner.databases.setIamPolicy",
    "spanner.databases.getIamPolicy",

  ]
}
