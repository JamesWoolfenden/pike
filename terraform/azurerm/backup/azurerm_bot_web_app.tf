resource "azurerm_bot_web_app" "pike_gen" {
  name                    = "example"
  location                = "global"
  resource_group_name     = "pike"
  sku                     = "F0"
  microsoft_app_id        = "pike"
  microsoft_app_type      = "SingleTenant"
  microsoft_app_tenant_id = "pike"
}
