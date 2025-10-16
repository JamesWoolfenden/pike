
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "compute.globalOperations.get",
    "compute.healthChecks.create",
    "compute.healthChecks.delete",
    "compute.healthChecks.get",
    "compute.healthChecks.update",
    "compute.healthChecks.useReadOnly",
    "compute.httpHealthChecks.useReadOnly",
    "compute.httpsHealthChecks.useReadOnly",
    "compute.regionBackendServices.create",
    "compute.regionBackendServices.delete",
    "compute.regionBackendServices.get",
    "compute.regionBackendServices.update",
    "compute.regionHealthChecks.useReadOnly",
    "resourcemanager.projects.get",
    "storage.buckets.get",

    "iap.webServices.getSettings",
    "iap.webServices.updateSettings"
  ]
}
