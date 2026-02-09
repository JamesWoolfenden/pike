
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    //google_vmwareengine_datastore -- data
    "vmwareengine.datastores.get",

    //google_vmwareengine_datastore
    "vmwareengine.datastores.get",
    "vmwareengine.datastores.create",
    "vmwareengine.datastores.delete",
    "vmwareengine.datastores.update"

  ]
}
