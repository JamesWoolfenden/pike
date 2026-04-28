resource "azurerm_api_management_user" "pike_gen" {
  user_id             = "5931a75ae4bbd512288c680b"
  api_management_name = "pike"
  resource_group_name = "pike"
  first_name          = "Example"
  last_name           = "User"
  email               = "tom+tfdev@hashicorp.com"
  state               = "active"
}
