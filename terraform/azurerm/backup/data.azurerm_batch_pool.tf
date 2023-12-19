data "azurerm_batch_pool" "pike" {
  name                = "pike"
  resource_group_name = "pike"
  account_name        = "pike"
}
