resource "azurerm_dev_center_project_pool" "pike_gen" {
  name                                    = "example-dcpl"
  location                                = "pike"
  dev_center_project_id                   = "pike"
  dev_box_definition_name                 = "pike"
  local_administrator_enabled             = true
  dev_center_attached_network_name        = "pike"
  stop_on_disconnect_grace_period_minutes = 60
}
