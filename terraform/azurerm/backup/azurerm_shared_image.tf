resource "azurerm_shared_image" "pike" {
  name                = "pike-image"
  gallery_name        = azurerm_shared_image_gallery.pike.name
  resource_group_name = "pike"
  location            = "uksouth"
  os_type             = "Linux"

  identifier {
    publisher = "PikePublisher"
    offer     = "PikeOffer"
    sku       = "PikeSku"
  }

  tags = {
    pike = "permissions"
  }
}
