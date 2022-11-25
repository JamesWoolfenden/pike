resource "azurerm_managed_disk" "pike" {
  name                 = "pike"
  location             = "uksouth"
  resource_group_name  = "pike"
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "1"

  tags = {
    pike    = "permissions"
    another = "one"
  }
}
