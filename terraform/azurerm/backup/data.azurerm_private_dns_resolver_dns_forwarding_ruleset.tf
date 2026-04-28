data "azurerm_private_dns_resolver_dns_forwarding_ruleset" "pike_gen" {
  name                = "example-ruleset"
  resource_group_name = "example-ruleset-resourcegroup"
}
