resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_service_directory_namespace
    "servicedirectory.namespaces.create",
    "servicedirectory.namespaces.get",
    "servicedirectory.namespaces.delete",
    "servicedirectory.namespaces.update",

    //google_service_directory_service
    "servicedirectory.services.create",
    "servicedirectory.services.get",
    "servicedirectory.services.delete",
    "servicedirectory.services.update",

    //google_service_directory_endpoint
    "servicedirectory.endpoints.create",
    "servicedirectory.endpoints.get",
    "servicedirectory.endpoints.delete",
    "servicedirectory.endpoints.update",

    //google_service_directory_namespace_iam_policy, google_service_directory_namespace_iam_binding
    //google_service_directory_namespace_iam_binding
    "servicedirectory.namespaces.getIamPolicy",
    "servicedirectory.namespaces.setIamPolicy",
    //google_service_directory_service_iam_policy, google_service_directory_namespace_iam_policy,
    //google_service_directory_service_iam_member, google_service_directory_service_iam_binding,
    "servicedirectory.services.getIamPolicy",
    "servicedirectory.services.setIamPolicy",

  ]
}
