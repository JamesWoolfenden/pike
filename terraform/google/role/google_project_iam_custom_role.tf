
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "cloudsql.instances.create",
    "cloudsql.instances.delete",
    "cloudsql.instances.get",
    "cloudsql.instances.update",
    "cloudsql.sslCerts.create",
    "cloudsql.sslCerts.delete",
    "cloudsql.sslCerts.get",
    "cloudsql.users.create",
    "cloudsql.users.delete",
    "cloudsql.users.list",
    "cloudsql.users.update",
    "compute.globalOperations.get",
    "compute.networks.create",
    "compute.networks.delete",
    "compute.networks.get",
    "dataplex.assets.create",
    "dataplex.assets.delete",
    "dataplex.assets.get",
    "dataplex.assets.update",
    "dataplex.entries.create",
    "dataplex.entries.delete",
    "dataplex.entries.get",
    "dataplex.entries.update",
    "dataplex.entryGroups.create",
    "dataplex.entryGroups.delete",
    "dataplex.entryGroups.get",
    "dataplex.entryGroups.update",
    "dataplex.entryTypes.create",
    "dataplex.entryTypes.delete",
    "dataplex.entryTypes.get",
    "dataplex.entryTypes.update",
    "dataplex.lakes.create",
    "dataplex.lakes.delete",
    "dataplex.lakes.get",
    "dataplex.lakes.update",
    "dataplex.zones.create",
    "dataplex.zones.delete",
    "dataplex.zones.get",
    "dataplex.zones.update",
    "resourcemanager.projects.get",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.update",

    //google_database_migration_service_connection_profile
    "datamigration.connectionprofiles.create",
    "datamigration.connectionprofiles.delete",
    "datamigration.connectionprofiles.get",
    "datamigration.connectionprofiles.update",
    "datamigration.operations.get",

    //google_database_migration_service_private_connection
    "datamigration.privateconnections.get",
    "datamigration.privateconnections.delete",
    "datamigration.privateconnections.create",
    "datamigration.operations.get",

    //google_data_pipeline_pipeline


  ]
}
