resource "azurerm_netapp_account_encryption" "pike_gen" {
  netapp_account_id = "pike"

  user_assigned_identity_id = "pike"

  encryption_key = "pike"

  # Optional: For cross-tenant key vault access scenarios
  federated_client_id = "pike"
}
