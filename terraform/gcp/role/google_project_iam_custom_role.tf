resource "google_project_iam_custom_role" "pike" {
  project     = "pike-361314"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "compute.subnetworks.create",
    "compute.networks.updatePolicy",
    "compute.subnetworks.get",
    "compute.subnetworks.delete"
  ]
}
