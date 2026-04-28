resource "azurerm_monitor_data_collection_endpoint" "pike_gen" {
  name                          = "example-mdce"
  resource_group_name           = "pike"
  location                      = "pike"
  kind                          = "Windows"
  public_network_access_enabled = true
  description                   = "monitor_data_collection_endpoint example"
  tags = {
    foo = "bar"
  }
}
