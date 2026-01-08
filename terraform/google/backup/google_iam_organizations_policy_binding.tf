# resource "time_sleep" "wait_60_seconds" {
#   create_duration = "60s"
#   depends_on = [google_iam_principal_access_boundary_policy.pab_policy]
# }

resource "google_iam_organizations_policy_binding" "binding-for-all-org-principals" {

  organization      = "123456789"
  location          = "global"
  display_name      = "binding for all principals in the Organization"
  policy_kind       = "PRINCIPAL_ACCESS_BOUNDARY"
  policy_binding_id = "binding-for-all-org-principals"
  policy            = "{}"
  target {
    principal_set = "//cloudresourcemanager.googleapis.com/organizations/123456789"
  }
}
