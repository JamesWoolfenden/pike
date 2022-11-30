resource "azurerm_redis_cache" "pike" {
  name                = "pike"
  location            = "uksouth"
  resource_group_name = "pike"
  capacity            = 2
  family              = "C"
  sku_name            = "Standard"
  enable_non_ssl_port = false
  minimum_tls_version = "1.2"

  redis_configuration {
  }
}
