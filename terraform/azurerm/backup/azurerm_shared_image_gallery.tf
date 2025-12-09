resource "azurerm_shared_image_gallery" "pike" {
  name                = "pike_gallery"
  resource_group_name = "pike"
  location            = "uksouth"
  description         = "Pike test gallery"

  tags = {
    pike = "permissions"
  }
}
