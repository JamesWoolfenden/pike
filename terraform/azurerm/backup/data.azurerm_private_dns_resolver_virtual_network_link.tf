data "azurerm_private_dns_resolver_virtual_network_link" "pike_gen" {
  name                      = "example-link"
  dns_forwarding_ruleset_id = "example-dns-forwarding-ruleset-id"
}
