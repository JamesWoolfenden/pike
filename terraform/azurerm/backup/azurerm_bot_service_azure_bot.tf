resource "azurerm_bot_service_azure_bot" "pike_gen" {
  name                    = "exampleazurebot"
  resource_group_name     = "pike"
  location                = "global"
  microsoft_app_id        = "pike"
  microsoft_app_type      = "SingleTenant"
  microsoft_app_tenant_id = "pike"
  sku                     = "F0"

  endpoint                              = "https://example.com"
  developer_app_insights_application_id = "pike"

  tags = {
    environment = "test"
  }
}
