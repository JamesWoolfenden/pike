resource "azurerm_api_management_product_policy" "pike_gen" {
  product_id          = "pike"
  api_management_name = "pike"
  resource_group_name = "pike"

  xml_content = <<XML
<policies>
  <inbound>
    <find-and-replace from="xyz" to="abc" />
  </inbound>
</policies>
XML

}
