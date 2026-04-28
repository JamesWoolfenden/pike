resource "azurerm_logic_app_integration_account_session" "pike_gen" {
  name                     = "example-ias"
  resource_group_name      = "pike"
  integration_account_name = "pike"

  content = <<CONTENT
 {
       "controlNumber": "1234"
    }
  CONTENT
}
