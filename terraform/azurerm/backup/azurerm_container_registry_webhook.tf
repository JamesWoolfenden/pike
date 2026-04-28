resource "azurerm_container_registry_webhook" "pike_gen" {
  name                = "mywebhook"
  resource_group_name = "pike"
  registry_name       = "pike"
  location            = "pike"

  service_uri = "https://mywebhookreceiver.example/mytag"
  status      = "enabled"
  scope       = "mytag:*"
  actions     = ["push"]
  custom_headers = {
    "Content-Type" = "application/json"
  }
}
