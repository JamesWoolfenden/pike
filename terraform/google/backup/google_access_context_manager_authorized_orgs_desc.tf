resource "google_access_context_manager_authorized_orgs_desc" "authorized-orgs-desc" {
  parent                  = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}"
  name                    = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}/authorizedOrgsDescs/fakeDescName"
  authorization_type      = "AUTHORIZATION_TYPE_TRUST"
  asset_type              = "ASSET_TYPE_CREDENTIAL_STRENGTH"
  authorization_direction = "AUTHORIZATION_DIRECTION_TO"
  orgs                    = ["organizations/12345", "organizations/98765"]
}

resource "google_access_context_manager_access_policy" "test-access" {
  parent = "organizations/1231231231"
  title  = "my policy"
}
