resource "azurerm_lb_outbound_rule" "pike_gen" {
  name                    = "OutboundRule"
  loadbalancer_id         = "pike"
  protocol                = "Tcp"
  backend_address_pool_id = "pike"

  frontend_ip_configuration {
    name = "PublicIPAddress"
  }
}
