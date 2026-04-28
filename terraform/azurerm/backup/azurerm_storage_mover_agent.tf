resource "azurerm_storage_mover_agent" "pike_gen" {
  name                     = "example-sa"
  storage_mover_id         = "pike"
  arc_virtual_machine_id   = "${azurerm_resource_group.example.id}/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"
  arc_virtual_machine_uuid = "3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"
  description              = "Example Agent Description"
}
