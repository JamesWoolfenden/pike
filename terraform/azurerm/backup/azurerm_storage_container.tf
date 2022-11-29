resource "azurerm_storage_container" "pike" {
  name                  = "vhds"
  storage_account_name  = "pike"
  container_access_type = "private"
}

output "container" {
  value = azurerm_storage_container.pike
}
