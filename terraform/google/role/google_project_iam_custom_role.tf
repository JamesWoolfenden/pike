
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "tpu.runtimeversions.list",
    "storageinsights.reportConfigs.get",
    "storage.buckets.get",
    "storage.buckets.delete",
    "storage.buckets.getIamPolicy",
    "storageinsights.reportConfigs.delete",
    "storage.buckets.setIamPolicy",

    //google_tpu_v2_vm google_tpu_v2_queued_resource
    "tpu.nodes.create",
    "tpu.nodes.get",
    "tpu.nodes.update",
    "tpu.nodes.delete"
  ]
}
