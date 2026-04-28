resource "azurerm_elastic_san" "pike_gen" {
  name                 = "example"
  resource_group_name  = "pike"
  location             = "pike"
  base_size_in_tib     = 1
  extended_size_in_tib = 2
  sku {
    name = "example-value"
  }
}
