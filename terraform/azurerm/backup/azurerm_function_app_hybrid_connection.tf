resource "azurerm_function_app_hybrid_connection" "pike_gen" {
  function_app_id = "pike"
  relay_id        = "pike"
  hostname        = "myhostname.example"
  port            = 8081
}
