resource "azurerm_iothub_enrichment" "pike_gen" {
  resource_group_name = "pike"
  iothub_name         = "pike"
  key                 = "example"

  value          = "my value"
  endpoint_names = ["pike"]
}
