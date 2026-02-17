
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "compute.disks.create",
    "compute.globalOperations.get",
    "compute.healthChecks.create",
    "compute.healthChecks.delete",
    "compute.healthChecks.get",
    "compute.healthChecks.update",
    "compute.instanceTemplates.create",
    "compute.instanceTemplates.delete",
    "compute.instanceTemplates.get",
    "compute.instances.create",
    "compute.instances.delete",
    "compute.instances.get",
    "compute.instances.setLabels",
    "compute.instances.setMetadata",
    "compute.networks.get",
    "compute.subnetworks.use",
    "compute.subnetworks.useExternalIp",
    "compute.zones.get",
    "storage.buckets.get",
    "compute.networks.list",
    "compute.disks.get",
    "compute.disks.delete",
    "compute.disks.setLabels",

    //google_compute_snapshot_settings


    //google_compute_snapshot

  ]
}
