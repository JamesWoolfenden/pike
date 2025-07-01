resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.update",

    "compute.globalNetworkEndpointGroups.delete",
    "compute.globalNetworkEndpointGroups.get",
    "compute.globalNetworkEndpointGroups.create",
    # "compute.globalNetworkEndpointGroups.update",
    "compute.globalOperations.get",

    "compute.regionNetworkEndpointGroups.create",
    "compute.regionNetworkEndpointGroups.get",
    "compute.regionNetworkEndpointGroups.delete",
    "compute.regionOperations.get",

    "compute.networkEndpointGroups.create",
    "compute.networkEndpointGroups.get",
    "compute.networkEndpointGroups.delete",
    //"compute.operations.get"

  ]
}
