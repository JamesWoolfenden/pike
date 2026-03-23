
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "compute.globalOperations.get",
    "compute.networks.create",
    "compute.networks.delete",
    "compute.networks.get",
    "compute.networks.updatePolicy",
    "compute.subnetworks.create",
    "compute.subnetworks.delete",
    "compute.subnetworks.get",
    "compute.subnetworks.update",
    "storage.buckets.get",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.update",
    "storage.objects.get",
    "storage.objects.list",
    "storage.objects.create",
    "storage.objects.delete",

    //Permissions to CREATE resources for testing data sources
    "logging.sinks.create",
    "logging.sinks.delete",
    "logging.sinks.update",
    "storage.managedFolders.create",
    "storage.managedFolders.delete",
    "appengine.applications.get",
    "datastore.databases.create",
    "datastore.databases.delete",
    "datastore.databases.get",
    "datastore.databases.update",
    "datastore.operations.get",
    "datastore.entities.create",
    "datastore.entities.delete",
    "datastore.entities.update",







    "biglake.tables.getIamPolicy"
  ]
}
