resource "azurerm_lb_nat_pool" "pike_gen" {
  resource_group_name            = "pike"
  loadbalancer_id                = "pike"
  name                           = "SampleApplicationPool"
  protocol                       = "Tcp"
  frontend_port_start            = 80
  frontend_port_end              = 81
  backend_port                   = 8080
  frontend_ip_configuration_name = "PublicIPAddress"
}
