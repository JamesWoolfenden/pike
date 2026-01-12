
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "storage.buckets.get",

    "compute.images.get",
    "compute.disks.create",
    "compute.disks.get",
    "compute.disks.update",
    "compute.disks.delete",
    "compute.disks.setLabels",
    "compute.zoneOperations.get"


  ]
}
