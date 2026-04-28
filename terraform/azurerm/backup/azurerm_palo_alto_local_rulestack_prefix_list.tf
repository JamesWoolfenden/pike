resource "azurerm_palo_alto_local_rulestack_prefix_list" "pike_gen" {
  name         = "example"
  rulestack_id = "pike"
  prefix_list  = ["10.0.1.0/24"]
}
