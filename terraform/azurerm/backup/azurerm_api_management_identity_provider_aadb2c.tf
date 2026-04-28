resource "azurerm_api_management_identity_provider_aadb2c" "pike_gen" {
  resource_group_name = "pike"
  api_management_name = "pike"
  client_id           = "pike"
  allowed_tenant      = "myb2ctenant.onmicrosoft.com"
  signin_tenant       = "myb2ctenant.onmicrosoft.com"
  authority           = "myb2ctenant.b2clogin.com"
  signin_policy       = "B2C_1_Login"
  signup_policy       = "B2C_1_Signup"

  depends_on = ["pike"]
}
