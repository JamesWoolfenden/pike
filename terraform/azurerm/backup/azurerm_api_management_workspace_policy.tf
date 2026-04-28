resource "azurerm_api_management_workspace_policy" "pike_gen" {
  api_management_workspace_id = "pike"
  xml_content                 = <<XML
<policies>
  <inbound>
    <find-and-replace from="abc" to="xyz" />
  </inbound>
</policies>
XML
}
