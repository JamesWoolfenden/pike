resource "azurerm_cognitive_deployment" "pike_gen" {
  name                 = "example-cd"
  cognitive_account_id = "pike"

  model {
    format  = "OpenAI"
    name    = "text-curie-001"
    version = "1"
  }

  sku {
    name = "Standard"
  }
}
