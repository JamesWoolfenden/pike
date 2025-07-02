
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "clientauthconfig.clients.get",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.update",
    "firebase.projects.create",
    "compute.globalNetworkEndpointGroups.delete",
    "compute.globalNetworkEndpointGroups.get",
    "compute.globalNetworkEndpointGroups.create",
    "compute.globalOperations.get",
    "compute.regionNetworkEndpointGroups.create",
    "compute.regionNetworkEndpointGroups.get",
    "compute.regionNetworkEndpointGroups.delete",
    "compute.regionOperations.get",
    //"compute.operations.get"
  ]
}
