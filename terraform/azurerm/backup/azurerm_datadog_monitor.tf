resource "azurerm_datadog_monitor" "pike_gen" {
  name                = "example-monitor"
  resource_group_name = "pike"
  location            = "pike"
  datadog_organization {
    application_key = "XXXX"
  }
  user {
    name  = "Example"
    email = "abc@xyz.com"
  }
  sku_name = "Linked"
  identity {
    type = "SystemAssigned"
  }
}
