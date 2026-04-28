resource "azurerm_maintenance_assignment_dynamic_scope" "pike_gen" {
  name                         = "example"
  maintenance_configuration_id = "pike"

  filter {
    locations       = ["West Europe"]
    os_types        = ["Windows"]
    resource_groups = ["pike"]
    resource_types  = ["Microsoft.Compute/virtualMachines"]
    tag_filter      = "Any"
    tags {
      tag    = "foo"
      values = ["barbar"]
    }
  }
}
