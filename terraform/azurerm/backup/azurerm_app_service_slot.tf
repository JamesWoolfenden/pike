resource "azurerm_app_service_slot" "pike_gen" {
  name                = "pike"
  app_service_name    = "pike"
  location            = "pike"
  resource_group_name = "pike"
  app_service_plan_id = "pike"

  site_config {
    dotnet_framework_version = "v4.0"
  }

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }
}
