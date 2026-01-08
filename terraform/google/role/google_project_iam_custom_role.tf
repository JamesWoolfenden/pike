
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "resourcemanager.projects.get",
    "appengine.applications.get",
    "datastore.operations.get",

    //docs
    "datastore.entities.create",
    "datastore.entities.delete",
    "datastore.entities.update",
    "datastore.entities.get",
    "datastore.namespaces.get"
  ]
}
