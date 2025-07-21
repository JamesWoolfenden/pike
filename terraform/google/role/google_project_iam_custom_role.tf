
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "resourcemanager.projects.get",
    //      "billing.resourceAssociations.delete",
    "resourcemanager.projects.deleteBillingAssignment",
    "resourcemanager.projects.createBillingAssignment"

  ]
}
