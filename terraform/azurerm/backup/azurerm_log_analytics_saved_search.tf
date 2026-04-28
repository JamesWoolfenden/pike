resource "azurerm_log_analytics_saved_search" "pike_gen" {
  name                       = "exampleSavedSearch"
  log_analytics_workspace_id = "pike"

  category     = "exampleCategory"
  display_name = "exampleDisplayName"
  query        = "exampleQuery"
}
