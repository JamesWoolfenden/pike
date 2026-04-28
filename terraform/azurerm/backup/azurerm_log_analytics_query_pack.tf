resource "azurerm_log_analytics_query_pack" "pike_gen" {
  name                = "example-laqp"
  resource_group_name = "pike"
  location            = "pike"
}
