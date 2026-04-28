resource "azurerm_network_manager_management_group_connection" "pike_gen" {
  name                = "example-nmmgc"
  management_group_id = "pike"
  network_manager_id  = "pike"
  description         = "example"
  depends_on          = ["pike"]
}
