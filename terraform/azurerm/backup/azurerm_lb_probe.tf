resource "azurerm_lb_probe" "pike_gen" {
  loadbalancer_id = "pike"
  name            = "ssh-running-probe"
  port            = 22
}
