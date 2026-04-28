resource "azurerm_ai_foundry_project" "pike_gen" {
  name               = "example"
  location           = "pike"
  ai_services_hub_id = "pike"
}
