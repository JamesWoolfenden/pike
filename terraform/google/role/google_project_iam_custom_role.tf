
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "notebooks.environments.create",
    "notebooks.environments.get",
    // "notebooks.environments.update",
    "notebooks.environments.delete",

    "notebooks.instances.create",
    "notebooks.instances.delete",
    "notebooks.instances.update",
    "notebooks.instances.get",

    "notebooks.runtimes.create",
    "notebooks.runtimes.delete",
    "notebooks.runtimes.get",
    "notebooks.runtimes.update"

  ]
}
