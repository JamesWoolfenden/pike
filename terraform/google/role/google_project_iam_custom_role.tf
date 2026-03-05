
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

    //google_storage_managed_folder_iam_policy
    "storage.managedFolders.get",
    "storage.managedFolders.getIamPolicy",

    //google_cloud_asset_search_all_resources
    "cloudasset.assets.searchAllResources",

    //google_firestore_document
    "datastore.entities.get",

    //google_iam_testable_permissions
    "iam.roles.list",

    //google_logging_sink
    "logging.sinks.get",

    //google_network_management_connectivity_test_run
    "networkmanagement.connectivitytests.get",

    //google_artifact_registry_package
    "artifactregistry.packages.get",

    //google_artifact_registry_tag
    "artifactregistry.tags.get",

    //google_backup_dr_backup_plan
    "backupdr.backupPlans.get",

    //google_backup_dr_management_server
    "backupdr.managementServers.list",

    //google_compute_region_ssl_policy
    "compute.regionSslPolicies.get",

    //google_vertex_ai_index
    "aiplatform.indexes.get",

    //google_network_management_connectivity_tests
    "networkmanagement.connectivitytests.list"
  ]
}
