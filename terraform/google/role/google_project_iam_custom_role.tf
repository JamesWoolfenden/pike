
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [

    //google_artifact_registry_python_packages
    "artifactregistry.pythonpackages.list",

    "compute.reservationBlocks.get",

    "compute.reservationSubBlocks.get"

  ]
}
