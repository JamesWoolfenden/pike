
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    "clientauthconfig.clients.get",
    "firebase.clients.create",
    "firebase.clients.delete",
    "firebase.clients.get",
    "firebase.clients.update",
    "firebase.projects.create",
    "firebase.projects.delete",
    "firebase.projects.get",
    "firebase.projects.update",
    "resourcemanager.projects.delete",
    "resourcemanager.projects.get",
    "resourcemanager.projects.update",
    "serviceusage.services.enable",
    "serviceusage.services.get",
    "storage.buckets.get"
  ]
}
