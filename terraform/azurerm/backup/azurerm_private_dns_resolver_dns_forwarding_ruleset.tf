resource "azurerm_private_dns_resolver_dns_forwarding_ruleset" "pike_gen" {
  name                                       = "example-ruleset"
  resource_group_name                        = "pike"
  location                                   = "pike"
  private_dns_resolver_outbound_endpoint_ids = ["pike"]
  tags = {
    key = "value"
  }
}
