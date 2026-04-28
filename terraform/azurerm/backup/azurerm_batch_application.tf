resource "azurerm_batch_application" "pike_gen" {
  name                = "example-batch-application"
  resource_group_name = "pike"
  account_name        = "pike"
}
