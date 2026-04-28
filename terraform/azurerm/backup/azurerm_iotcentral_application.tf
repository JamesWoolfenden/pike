resource "azurerm_iotcentral_application" "pike_gen" {
  name                = "example-iotcentral-app"
  resource_group_name = "pike"
  location            = "pike"
  sub_domain          = "example-iotcentral-app-subdomain"

  display_name = "example-iotcentral-app-display-name"
  sku          = "ST1"
  template     = "iotc-default@1.0.0"

  tags = {
    Foo = "Bar"
  }
}
