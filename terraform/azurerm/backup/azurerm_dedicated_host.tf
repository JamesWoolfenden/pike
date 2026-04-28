resource "azurerm_dedicated_host" "pike_gen" {
  name                    = "example-host"
  location                = "pike"
  dedicated_host_group_id = "pike"
  sku_name                = "DSv3-Type3"
  platform_fault_domain   = 1
}
