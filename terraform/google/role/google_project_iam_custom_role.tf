
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    //google_compute_region_backend_bucket
    "compute.regionBackendBuckets.create",
    "compute.regionBackendBuckets.delete",
    "compute.regionBackendBuckets.get",
    "compute.regionBackendBuckets.update",

    //google_compute_region_backend_bucket_iam_policy
    "compute.regionBackendBuckets.getIamPolicy",
    "compute.regionBackendBuckets.setIamPolicy",

    //storage bucket dependency
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",

    //google_biglake_iceberg_catalog_iam_policy
    "biglake.catalogs.getIamPolicy",
    "biglake.catalogs.setIamPolicy",

    //google_biglake_iceberg_namespace_iam_policy
    "biglake.namespaces.getIamPolicy",
    "biglake.namespaces.setIamPolicy",

    ///google_biglake_iceberg_namespace
    "biglake.namespaces.create",
    "biglake.namespaces.delete",
    "biglake.namespaces.get",
    "biglake.namespaces.update",

    //google_backup_dr_data_sources
    "backupdr.bvdataSources.list"
  ]
}
