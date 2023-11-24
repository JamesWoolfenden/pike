resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_sourcerepo_repository
    "source.repos.get",
    //google_sourcerepo_repository_iam_policy
    "source.repos.getIamPolicy",
    //
    "cloudsql.backupRuns.get",
    "cloudsql.backupRuns.list",

    //google_sql_database
    "cloudsql.databases.get",
    //google_sql_database_instance
    "cloudsql.instances.get",
    "cloudsql.instances.list",

    "cloudsql.databases.list",

    //notsure
    "cloudsql.sslCerts.get",
    "cloudsql.sslCerts.list"
  ]
}
