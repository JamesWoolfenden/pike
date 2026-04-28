resource "azurerm_system_center_virtual_machine_manager_virtual_machine_instance" "pike_gen" {
  scoped_resource_id = "pike"
  custom_location_id = "pike"

  infrastructure {
    system_center_virtual_machine_manager_cloud_id                  = "pike"
    system_center_virtual_machine_manager_template_id               = "pike"
    system_center_virtual_machine_manager_virtual_machine_server_id = "pike"
  }

  operating_system {
    computer_name = "testComputer"
  }

  hardware {
    cpu_count    = 1
    memory_in_mb = 1024
  }

  lifecycle {
    // Service API always provisions a virtual disk with bus type IDE per Virtual Machine Template by default, so it has to be ignored
    ignore_changes = [storage_disk]
  }
}
