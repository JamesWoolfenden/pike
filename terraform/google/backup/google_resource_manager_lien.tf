resource "google_resource_manager_lien" "pike" {
  parent       = "projects/430943803039"
  restrictions = ["resourcemanager.projects.delete"]
  origin       = "pike-test"
  reason       = "Pike test lien - prevent deletion during testing"
}
