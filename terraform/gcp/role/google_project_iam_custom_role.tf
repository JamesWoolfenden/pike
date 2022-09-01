resource "google_project_iam_custom_role" "pike" {
  project     = "examplea"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "storage.objects.create",
    "storage.objects.delete",

    "storage.objects.get"
  ]
}
