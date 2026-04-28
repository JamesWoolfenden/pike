resource "azurerm_federated_identity_credential" "pike_gen" {
  name                      = "example"
  audience                  = ["foo"]
  issuer                    = "https://foo"
  user_assigned_identity_id = "pike"
  subject                   = "foo"
}
