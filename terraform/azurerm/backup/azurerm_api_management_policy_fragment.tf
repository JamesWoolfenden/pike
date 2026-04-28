resource "azurerm_api_management_policy_fragment" "pike_gen" {
  api_management_id = "pike"
  name              = "example-policy-fragment"
  format            = "xml"
  value             = file("policy-fragment-1.xml")
}
