resource "google_project_iam_custom_role" "pike" {
  project     = "pike-361314"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "compute.networks.get",
    "compute.networks.list",
    "resourcemanager.projects.get",
    "servicenetworking.services.get",
    "servicenetworking.services.addPeering",
    "compute.networks.removePeering",
    "compute.serviceAttachments.get",
    "compute.globalAddresses.create",
    "compute.globalAddresses.delete",
    "compute.globalAddresses.createInternal",
    "compute.globalAddresses.deleteInternal",
    "compute.networks.use",
    "compute.globalAddresses.get"
  ]
}
