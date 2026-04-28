resource "azurerm_dev_center_project_environment_type" "pike_gen" {
  name                  = "example-et"
  location              = "pike"
  dev_center_project_id = "pike"
  deployment_target_id  = "/subscriptions/${data.azurerm_client_config.current.subscription_id}"

  identity {
    type = "SystemAssigned"
  }
}
