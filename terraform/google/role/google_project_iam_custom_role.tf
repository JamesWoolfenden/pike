
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "aiplatform.indexEndpoints.create",
    "aiplatform.indexes.create",
    "compute.networks.get",
    "iam.serviceAccounts.create",
    "iam.serviceAccounts.delete",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.update",
    "resourcemanager.projects.get",
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.update",
    "storage.objects.create",
    "storage.objects.delete",
    "storage.objects.get",
    "storage.objects.list",

    //new
    //google_vertex_ai_endpoint_with_model_garden_deployment
    "aiplatform.endpoints.create",
    "aiplatform.endpoints.delete",
    "aiplatform.endpoints.get",
    "aiplatform.endpoints.update",
    "aiplatform.endpoints.deploy",
    "aiplatform.models.upload"
  ]
}
