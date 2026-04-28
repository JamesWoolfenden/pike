resource "azurerm_databox_edge_device" "pike_gen" {
  name                = "example-device"
  resource_group_name = "pike"
  location            = "pike"

  sku_name = "EdgeP_Base-Standard"
}
