data "azurerm_lb_outbound_rule" "pike_gen" {
  name            = "existing_lb_outbound_rule"
  loadbalancer_id = "existing_load_balancer_id"
}
