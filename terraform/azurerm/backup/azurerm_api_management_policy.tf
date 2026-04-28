resource "azurerm_api_management_policy" "pike_gen" {
  api_management_id = "pike"
  xml_content       = file("example.xml")
}
