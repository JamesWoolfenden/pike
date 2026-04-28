data "azurerm_lb_rule" "pike_gen" {
  name                = "first"
  resource_group_name = "example-resources"
  loadbalancer_id     = "pike"
}
