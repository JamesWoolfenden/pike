resource "azurerm_cognitive_account_project" "pike_gen" {
  name                 = "example-project"
  cognitive_account_id = "pike"
  location             = "pike"
  description          = "Example cognitive services project"
  display_name         = "Example Project"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "test"
  }
}
