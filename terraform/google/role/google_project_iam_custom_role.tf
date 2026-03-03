
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "cloudkms.cryptoKeys.create",
    "cloudkms.cryptoKeys.get",
    "cloudkms.keyRings.create",
    "cloudkms.keyRings.get",
    "compute.commitments.create",
    "compute.commitments.get",
    "compute.commitments.update",
    "compute.healthChecks.use",
    "compute.instanceGroupManagers.create",
    "compute.instanceGroupManagers.delete",
    "compute.instanceGroupManagers.get",
    "compute.instanceGroupManagers.update",
    "compute.instanceTemplates.create",
    "compute.instanceTemplates.delete",
    "compute.instanceTemplates.get",
    "compute.instanceTemplates.useReadOnly",
    "compute.networks.get",
    "logging.queries.create",
    "logging.queries.delete",
    "logging.queries.get",
    "logging.queries.update",
    "logging.queries.usePrivate",
    "netapp.kmsConfigs.create",
    "netapp.kmsConfigs.delete",
    "netapp.kmsConfigs.get",
    "netapp.kmsConfigs.update",
    "netapp.operations.get",
    "storage.buckets.get"
  ]
}
