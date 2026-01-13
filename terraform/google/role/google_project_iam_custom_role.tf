
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "storage.buckets.get",
    "compute.networks.get",

    "compute.disks.get",
    "compute.disks.delete",
    "compute.zoneOperations.get",
    //datasource
    //google_biglake_iceberg_catalog_iam_policy
    //none

    //google_artifact_registry_package
    //missing






  ]
}
