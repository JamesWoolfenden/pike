data "azurerm_shared_image" "pike" {
  resource_group_name = "pike"
  gallery_name        = "pike"
  name                = "pike"
}
