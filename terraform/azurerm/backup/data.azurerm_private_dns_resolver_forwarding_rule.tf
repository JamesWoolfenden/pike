data "azurerm_private_dns_resolver_forwarding_rule" "pike_gen" {
  name                      = "example-rule"
  dns_forwarding_ruleset_id = "example-forwarding-rulset-id"
}
