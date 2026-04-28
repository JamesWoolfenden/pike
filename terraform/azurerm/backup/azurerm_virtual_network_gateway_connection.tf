resource "azurerm_virtual_network_gateway_connection" "pike_gen" {
  name                = "onpremise"
  location            = "pike"
  resource_group_name = "pike"

  type                       = "IPsec"
  virtual_network_gateway_id = "pike"
  local_network_gateway_id   = "pike"

  shared_key = "4-v3ry-53cr37-1p53c-5h4r3d-k3y"
}
