data "azurerm_shared_image_versions" "pike" {
  resource_group_name = "pike"
  image_name          = "pike"
  gallery_name        = "pike"
}
