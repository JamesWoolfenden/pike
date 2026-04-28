resource "azurerm_dedicated_host_group" "pike_gen" {
  name                        = "example-dedicated-host-group"
  resource_group_name         = "pike"
  location                    = "pike"
  platform_fault_domain_count = 1
}
