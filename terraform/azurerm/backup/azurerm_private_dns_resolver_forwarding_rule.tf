resource "azurerm_private_dns_resolver_forwarding_rule" "pike_gen" {
  name                      = "example-rule"
  dns_forwarding_ruleset_id = "pike"
  domain_name               = "onprem.local."
  enabled                   = true
  target_dns_servers {
    ip_address = "10.10.0.1"
    port       = 53
  }
  metadata = {
    key = "value"
  }
}
