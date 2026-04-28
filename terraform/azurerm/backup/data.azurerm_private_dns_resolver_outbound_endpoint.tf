data "azurerm_private_dns_resolver_outbound_endpoint" "pike_gen" {
  name                    = "example-endpoint"
  private_dns_resolver_id = "example-private-dns-resolver-id"
}
