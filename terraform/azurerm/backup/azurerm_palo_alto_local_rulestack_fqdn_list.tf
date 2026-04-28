resource "azurerm_palo_alto_local_rulestack_fqdn_list" "pike_gen" {
  name         = "example"
  rulestack_id = "pike"

  fully_qualified_domain_names = ["contoso.com"]
}
