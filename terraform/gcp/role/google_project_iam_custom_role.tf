resource "google_project_iam_custom_role" "pike" {
  project     = "pike-361314"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "storage.buckets.create"
  ]
}
