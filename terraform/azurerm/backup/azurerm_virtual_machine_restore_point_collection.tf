resource "azurerm_virtual_machine_restore_point_collection" "pike_gen" {
  name                      = "example-collection"
  resource_group_name       = "pike"
  location                  = "pike"
  source_virtual_machine_id = "pike"
}
