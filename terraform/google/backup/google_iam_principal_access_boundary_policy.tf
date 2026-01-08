resource "google_iam_principal_access_boundary_policy" "pab_policy" {
  organization                        = "123456789"
  location                            = "global"
  display_name                        = "binding for all principals in the project"
  principal_access_boundary_policy_id = "my-pab-policy"
}
