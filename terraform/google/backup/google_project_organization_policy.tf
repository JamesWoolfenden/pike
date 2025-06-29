resource "google_project_organization_policy" "pike" {
  project    = "pike-plus"
  constraint = "compute.disableSerialPortAccess"

  boolean_policy {
    enforced = true
  }
}
