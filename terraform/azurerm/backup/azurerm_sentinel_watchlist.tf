resource "azurerm_sentinel_watchlist" "pike_gen" {
  name                       = "example-watchlist"
  log_analytics_workspace_id = "pike"
  display_name               = "example-wl"
  item_search_key            = "Key"
}
