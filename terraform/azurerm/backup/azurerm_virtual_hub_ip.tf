resource "azurerm_virtual_hub_ip" "pike_gen" {
  name                         = "example-vhubipconfig"
  virtual_hub_id               = "pike"
  private_ip_address           = "10.5.1.18"
  private_ip_allocation_method = "Static"
  public_ip_address_id         = "pike"
  subnet_id                    = "pike"
}
