resource "azurerm_virtual_machine_scale_set_standby_pool" "pike_gen" {
  name                                  = "example-spsvmp"
  resource_group_name                   = "pike"
  location                              = "West Europe"
  attached_virtual_machine_scale_set_id = "pike"
  virtual_machine_state                 = "Running"

  elasticity_profile {
    max_ready_capacity = 10
    min_ready_capacity = 5
  }

  tags = {
    key = "value"
  }
}
