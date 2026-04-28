resource "azurerm_vmware_private_cloud" "pike_gen" {
  name                = "example-vmware-private-cloud"
  resource_group_name = "pike"
  location            = "pike"
  sku_name            = "av36"

  management_cluster {
    size = 3
  }

  network_subnet_cidr         = "192.168.48.0/22"
  internet_connection_enabled = false
}
