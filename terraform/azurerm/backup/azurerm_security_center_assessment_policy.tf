resource "azurerm_security_center_assessment_policy" "pike_gen" {
  display_name = "Test Display Name"
  severity     = "Medium"
  description  = "Test Description"
}
