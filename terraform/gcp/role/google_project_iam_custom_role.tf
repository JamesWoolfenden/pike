resource "google_project_iam_custom_role" "pike" {
  project     = "examplea"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "compute.zones.get",
    "compute.instances.create",
    "compute.instances.get",
    "compute.disks.create",
    "compute.disks.create",
    "compute.subnetworks.use",
    "compute.subnetworks.useExternalIp",
    "compute.instances.setMetadata",
    "compute.instances.delete",
    "compute.instances.setTags"
  ]
}
