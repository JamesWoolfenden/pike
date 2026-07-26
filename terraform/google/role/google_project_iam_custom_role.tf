
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "compute.globalAddresses.create",
    "compute.globalAddresses.createInternal",
    "compute.globalAddresses.delete",
    "compute.globalAddresses.deleteInternal",
    "compute.globalAddresses.get",
    "compute.globalAddresses.setLabels",
    "compute.globalOperations.get",
    "compute.networks.create",
    "compute.networks.delete",
    "compute.networks.get",
    "compute.networks.removePeering",
    "compute.networks.use",
    "networkconnectivity.operations.get",
    "networkconnectivity.policyBasedRoutes.create",
    "networkconnectivity.policyBasedRoutes.delete",
    "networkconnectivity.policyBasedRoutes.get",
    "resourcemanager.projects.get",
    "servicenetworking.services.addPeering",
    "servicenetworking.services.createPeeredDnsDomain",
    "servicenetworking.services.deleteConnection",
    "servicenetworking.services.deletePeeredDnsDomain",
    "servicenetworking.services.get",
    "servicenetworking.services.listPeeredDnsDomains",
    "serviceusage.operations.get",
    "storage.buckets.get"
  ]
}
