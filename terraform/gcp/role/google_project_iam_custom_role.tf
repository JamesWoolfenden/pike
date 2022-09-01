resource "google_project_iam_custom_role" "pike" {
  project     = "examplea"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "storage.buckets.create",
    "storage.buckets.get",
    "storage.buckets.delete",

    "storage.buckets.update"
  ]
}
