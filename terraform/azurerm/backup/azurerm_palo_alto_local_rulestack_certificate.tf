resource "azurerm_palo_alto_local_rulestack_certificate" "pike_gen" {
  name         = "example"
  rulestack_id = "pike"
  self_signed  = true
}
