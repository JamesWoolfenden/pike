resource "azurerm_dev_test_lab" "pike_gen" {
  name                = "example-devtestlab"
  location            = "pike"
  resource_group_name = "pike"

  tags = {
    "Sydney" = "Australia"
  }
}
