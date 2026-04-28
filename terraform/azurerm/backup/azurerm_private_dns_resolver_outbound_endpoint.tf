resource "azurerm_private_dns_resolver_outbound_endpoint" "pike_gen" {
  name                    = "example-endpoint"
  private_dns_resolver_id = "pike"
  location                = "pike"
  subnet_id               = "pike"
  tags = {
    key = "value"
  }
}
