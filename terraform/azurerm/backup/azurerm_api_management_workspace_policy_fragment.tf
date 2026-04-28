resource "azurerm_api_management_workspace_policy_fragment" "pike_gen" {
  name                        = "example-policy-fragment"
  api_management_workspace_id = "pike"
  xml_format                  = "xml"
  xml_content                 = file("policy-fragment-1.xml")
}
