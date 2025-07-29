
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "resourcemanager.projects.get",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.getIamPolicy",
    "storage.buckets.setIamPolicy",
    "storage.buckets.update",
    "storageinsights.reportConfigs.create",


    "storage.buckets.get",
    "storage.objects.list",
    "storage.buckets.getObjectInsights",
    "storage.objects.create",
    "storage.buckets.get",
    "storageinsights.reportConfigs.get",
    "storageinsights.reportConfigs.create",
    "storageinsights.reportConfigs.delete",
    "storageinsights.reportConfigs.update"
  ]
}
