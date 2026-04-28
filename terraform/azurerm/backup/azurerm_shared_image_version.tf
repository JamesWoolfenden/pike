resource "azurerm_shared_image_version" "pike_gen" {
  name                = "0.0.1"
  gallery_name        = "pike"
  image_name          = "pike"
  resource_group_name = "pike"
  location            = "pike"
  managed_image_id    = "pike"

  target_region {
    name                   = "pike"
    regional_replica_count = 5
    storage_account_type   = "Standard_LRS"
  }
}
