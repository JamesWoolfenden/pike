resource "azurerm_workloads_sap_discovery_virtual_instance" "pike_gen" {
  name                              = "X01"
  resource_group_name               = "pike"
  location                          = "pike"
  environment                       = "NonProd"
  sap_product                       = "S4HANA"
  central_server_virtual_machine_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/csvm1"
  managed_storage_account_name      = "managedsa"

  identity {
    type = "UserAssigned"

    identity_ids = [
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai1"
    ]
  }

  lifecycle {
    ignore_changes = [managed_resource_group_name]
  }
}
