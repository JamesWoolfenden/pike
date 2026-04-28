resource "azurerm_automation_certificate" "pike_gen" {
  name                    = "certificate1"
  resource_group_name     = "pike"
  automation_account_name = "pike"

  description = "This is an example certificate"
  base64      = filebase64("certificate.pfx")
  exportable  = true
}
