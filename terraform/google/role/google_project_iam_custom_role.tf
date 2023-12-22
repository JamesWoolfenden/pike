resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //"composer.environments.list",
    "composer.environments.get",
    "composer.imageversions.list",

    //
    "cloudsql.instances.addServerCa",
    "cloudsql.instances.connect",
    "cloudsql.instances.export",
    "cloudsql.instances.failover",
    "cloudsql.instances.get",
    "cloudsql.instances.list",
    "cloudsql.instances.listServerCas",
    "cloudsql.instances.migrate",
    "cloudsql.instances.reencrypt",
    "cloudsql.instances.restart",
    "cloudsql.instances.rotateServerCa",
    "cloudsql.instances.truncateLog",
    "cloudsql.instances.update",
    "cloudsql.databases.create",
    "cloudsql.databases.get",
    "cloudsql.databases.list",
    "cloudsql.databases.update",
    "cloudsql.backupRuns.create",
    "cloudsql.backupRuns.get",
    "cloudsql.backupRuns.list",
    "cloudsql.sslCerts.get",
    "cloudsql.sslCerts.list",
    "cloudsql.users.list",

  ]
}
