resource "azurerm_monitor_private_link_scope" "pike_gen" {
  name                = "example-ampls"
  resource_group_name = "pike"

  ingestion_access_mode = "PrivateOnly"
  query_access_mode     = "Open"
}
