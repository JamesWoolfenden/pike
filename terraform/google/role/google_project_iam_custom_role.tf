resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "cloudbuild.builds.create",
    "cloudbuild.builds.update",
    "cloudbuild.builds.get"
  ]
}
