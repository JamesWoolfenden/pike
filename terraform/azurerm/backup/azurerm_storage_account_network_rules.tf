resource "azurerm_storage_account_network_rules" "pike" {
  default_action             = "Allow"
  ip_rules                   = ["127.0.0.1"]
  virtual_network_subnet_ids = ["/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.Network/virtualNetworks/pike/subnets/pike"]
  bypass                     = ["Metrics"]
  storage_account_id         = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Storage/storageAccounts/pike"
}
