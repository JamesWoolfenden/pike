data "google_privileged_access_manager_entitlement" "pike" {
  parent         = "projects/pike"
  location       = "global"
  entitlement_id = "my-entitlement"
  depends_on = [
    google_privileged_access_manager_entitlement.entitlement
  ]
}

output "google_privileged_access_manager_entitlement" {
  value = data.google_privileged_access_manager_entitlement.pike
}
