resource "azurerm_lb_rule" "pike_gen" {
  loadbalancer_id                = "pike"
  name                           = "LBRule"
  protocol                       = "Tcp"
  frontend_port                  = 3389
  backend_port                   = 3389
  frontend_ip_configuration_name = "PublicIPAddress"
}
