resource "azurerm_system_center_virtual_machine_manager_virtual_machine_instance_guest_agent" "pike_gen" {
  scoped_resource_id = "pike"
  username           = "Administrator"

  depends_on = ["pike"]
}
