data "azurerm_private_dns_resolver_inbound_endpoint" "pike_gen" {
  name                    = "example-drie"
  private_dns_resolver_id = "example-private-dns-resolver-id"
}
