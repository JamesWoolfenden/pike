resource "azurerm_virtual_hub_security_partner_provider" "pike_gen" {
  name                   = "example-spp"
  resource_group_name    = "pike"
  location               = "pike"
  virtual_hub_id         = "pike"
  security_provider_name = "IBoss"

  tags = {
    ENV = "Prod"
  }

  depends_on = ["pike"]
}
