resource "azurerm_relay_hybrid_connection" "pike_gen" {
  name                          = "acctestrnhc-%d"
  resource_group_name           = "pike"
  relay_namespace_name          = "pike"
  requires_client_authorization = false
  user_metadata                 = "testmetadata"
}
