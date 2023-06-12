data "azurerm_storage_account" "pike" {
  name                = "pike"
  resource_group_name = "pike"
}

output "account" {
  sensitive = true
  value     = data.azurerm_storage_account.pike
}
