resource "azurerm_lb_nat_rule" "pike_gen" {
  resource_group_name            = "pike"
  loadbalancer_id                = "pike"
  name                           = "RDPAccess"
  protocol                       = "Tcp"
  frontend_port                  = 3389
  backend_port                   = 3389
  frontend_ip_configuration_name = "PublicIPAddress"
}
