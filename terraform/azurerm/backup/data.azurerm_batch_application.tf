data "azurerm_batch_application" "pike" {
  resource_group_name = "pike"
  name                = "pike"
  account_name        = "pike"
}
