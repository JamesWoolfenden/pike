resource "azurerm_web_app_hybrid_connection" "pike_gen" {
  web_app_id = "pike"
  relay_id   = "pike"
  hostname   = "myhostname.example"
  port       = 8081
}
