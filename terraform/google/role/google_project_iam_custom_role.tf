resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "resourcemanager.projects.get",

    //google_pubsub_topic_iam_member
    "pubsub.topics.getIamPolicy",
    "pubsub.topics.setIamPolicy",

    //google_secret_manager_secret
    "secretmanager.secrets.create",
    "secretmanager.secrets.get",
    "secretmanager.secrets.update",
    "secretmanager.secrets.delete",

    //google_secret_manager_secret_iam_member
    "secretmanager.secrets.getIamPolicy",
    "secretmanager.secrets.setIamPolicy",

    //google_secret_manager_secret_version
    "secretmanager.versions.add",
    "secretmanager.versions.enable",
    "secretmanager.versions.get",
    "secretmanager.versions.access",
    "secretmanager.versions.destroy"
  ]
}
