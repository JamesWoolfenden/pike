resource "azurerm_elastic_cloud_elasticsearch" "pike_gen" {
  name                        = "example-elasticsearch"
  resource_group_name         = "pike"
  location                    = "pike"
  sku_name                    = "ess-consumption-2024_Monthly"
  elastic_cloud_email_address = "user@example.com"
}
