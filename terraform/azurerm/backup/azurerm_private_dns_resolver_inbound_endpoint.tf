resource "azurerm_private_dns_resolver_inbound_endpoint" "pike_gen" {
  name                    = "example-drie"
  private_dns_resolver_id = "pike"
  location                = "pike"
  ip_configurations {
    private_ip_allocation_method = "Dynamic"
    subnet_id                    = "pike"
  }
  tags = {
    key = "value"
  }
}
