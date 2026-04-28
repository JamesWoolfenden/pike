resource "azurerm_logic_app_action_http" "pike_gen" {
  name         = "webhook"
  logic_app_id = "pike"
  method       = "GET"
  uri          = "http://example.com/some-webhook"
}
