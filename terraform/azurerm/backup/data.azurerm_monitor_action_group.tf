data "azurerm_monitor_action_group" "pike_gen" {
  resource_group_name = "terraform-example-rg"
  name                = "tfex-actiongroup"
}
