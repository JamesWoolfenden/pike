resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_artifact_registry_repository
    "artifactregistry.repositories.create",
    "artifactregistry.repositories.get",
    "artifactregistry.repositories.update",
    "artifactregistry.repositories.delete",

    //google_artifact_registry_repository_iam_policy
    "artifactregistry.repositories.setIamPolicy",
    //google_artifact_registry_repository_iam_member, google_artifact_registry_repository_iam_binding
    "artifactregistry.repositories.getIamPolicy"
  ]
}
