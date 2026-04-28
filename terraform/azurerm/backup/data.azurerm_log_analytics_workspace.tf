data "azurerm_log_analytics_workspace" "pike_gen" {
  name                = "acctest-01"
  resource_group_name = "acctest"
}
