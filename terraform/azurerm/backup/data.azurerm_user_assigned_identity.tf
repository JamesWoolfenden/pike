data "azurerm_user_assigned_identity" "pike" {
  resource_group_name = "pike"
  name                = "pike"
}
