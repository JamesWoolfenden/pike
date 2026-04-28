resource "azurerm_sentinel_watchlist_item" "pike_gen" {
  name         = "0aac6fa5-223e-49cf-9bfd-3554dc9d2b76"
  watchlist_id = "pike"
  properties = {
    k1 = "v1"
    k2 = "v2"
  }
}
