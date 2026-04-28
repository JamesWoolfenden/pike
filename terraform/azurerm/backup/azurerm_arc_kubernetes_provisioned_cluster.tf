resource "azurerm_arc_kubernetes_provisioned_cluster" "pike_gen" {
  name                = "example-akpc"
  resource_group_name = "pike"
  location            = "pike"

  azure_active_directory {
    azure_rbac_enabled     = true
    admin_group_object_ids = ["pike"]
    tenant_id              = "pike"
  }

  identity {
    type = "SystemAssigned"
  }
}
