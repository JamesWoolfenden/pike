
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.update",
    "storage.buckets.get",


    "biglake.catalogs.create",
    "biglake.catalogs.delete",
    "biglake.catalogs.get",
    "biglake.namespaces.create",
    "biglake.namespaces.delete",
    "biglake.namespaces.get",
    "biglake.namespaces.update",
    "bigquery.datasets.create",
    "bigquery.datasets.delete",
    "bigquery.datasets.get",
    "bigquery.datasets.update",
    "bigquery.routines.create",
    "bigquery.routines.delete",
    "bigquery.routines.get",
    "bigquery.routines.update",
    "compute.firewallPolicies.create",
    "compute.firewallPolicies.delete",
    "compute.firewallPolicies.get",
    "compute.firewallPolicies.update",

    "compute.disks.create",
    "compute.disks.createSnapshot",
    "compute.disks.delete",
    "compute.disks.get",
    "compute.disks.setLabels",
    "compute.disks.update",
    "compute.disks.use",
    "compute.disks.useReadOnly",

    //permissions not valid will need to change these in pike if the permissions above work
    # "compute.regionDisks.create",
    # "compute.regionDisks.createSnapshot",
    # "compute.regionDisks.delete",
    # "compute.regionDisks.get",
    # "compute.regionDisks.setLabels",
    # "compute.regionDisks.update",
    # "compute.regionDisks.use",
    # "compute.regionDisks.useReadOnly",
    "dataplex.dataProducts.create",
    "dataplex.dataProducts.delete",
    "dataplex.dataProducts.get",
    "dataplex.dataProducts.update",
    "discoveryengine.dataStores.create",
    "discoveryengine.dataStores.delete",
    "discoveryengine.dataStores.get",
    "discoveryengine.dataStores.update",
    "discoveryengine.engines.create",
    "discoveryengine.engines.delete",
    "discoveryengine.engines.get",
    "discoveryengine.engines.update",
    "networkconnectivity.hubs.create",
    "networkconnectivity.hubs.delete",
    "networkconnectivity.hubs.get",
    "networkconnectivity.hubs.update",
    "networkconnectivity.operations.get",
    "storage.buckets.get",


    # google_compute_region_network_firewall_policy_iam_policy
    "compute.regionFirewallPolicies.getIamPolicy",
    "compute.regionFirewallPolicies.setIamPolicy",



    # google_iap_agent_registry_iam_binding
    # google_iap_agent_registry_iam_member
    # google_iap_agent_registry_iam_policy
    # google_iap_location_web_iam_binding
    # google_iap_location_web_iam_member
    # google_iap_location_web_iam_policy
    "iap.web.setIamPolicy",
    "iap.web.getIamPolicy",



    //google_discovery_engine_data_store
    "discoveryengine.dataStores.create",
    "discoveryengine.dataStores.get",
    "discoveryengine.dataStores.update",
    "discoveryengine.dataStores.delete",
  ]
}
