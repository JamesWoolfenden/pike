resource "azurerm_api_management_api" "pike_gen" {
  name                = "example-api"
  resource_group_name = "pike"
  api_management_name = "pike"
  revision            = "1"
  display_name        = "Example API"
  path                = "example"
  protocols           = ["https"]
  import {
    content_format = "swagger-link-json"
    content_value  = "https://raw.githubusercontent.com/hashicorp/terraform-provider-azurerm/refs/heads/main/internal/services/apimanagement/testdata/api_management_api_schema_swagger.json"
  }
}
