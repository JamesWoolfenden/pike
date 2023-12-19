data "azurerm_shared_image_version" "pike" {
  resource_group_name = "pike"
  image_name          = "pike"
  name                = "recent"
  gallery_name        = "pike"
}
