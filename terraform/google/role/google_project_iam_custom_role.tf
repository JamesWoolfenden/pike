
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    //google_network_security_address_groups
    "networksecurity.addressGroups.list",







    //google_dns_record_sets
    "dns.resourceRecordSets.get",

    //google_discovery_engine_data_store
    "discoveryengine.dataStores.get",


    //google_discovery_engine_data_stores
    "discoveryengine.dataStores.list",

    "storage.objects.get",
    "appengine.applications.get",
    "compute.globalOperations.get",
    "compute.networks.create",
    "compute.networks.delete",
    "compute.networks.get",
    "compute.networks.updatePolicy",
    "compute.subnetworks.create",
    "compute.subnetworks.delete",
    "compute.subnetworks.get",
    "compute.subnetworks.update",
    "datastore.databases.create",
    "datastore.databases.delete",
    "datastore.databases.get",
    "datastore.databases.getMetadata",
    "datastore.databases.update",
    "datastore.operations.get",
    "logging.sinks.create",
    "logging.sinks.delete",
    "logging.sinks.get",
    "logging.sinks.update",
    "resourcemanager.projects.get",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.update",
    # "storage.managedFolders.create",
    "storage.managedFolders.delete",
    "storage.managedFolders.get",
    # "storage.managedFolders.update"
  ]
}
