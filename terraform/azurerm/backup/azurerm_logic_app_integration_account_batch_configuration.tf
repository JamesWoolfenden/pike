resource "azurerm_logic_app_integration_account_batch_configuration" "pike_gen" {
  name                     = "exampleiabc"
  resource_group_name      = "pike"
  integration_account_name = "pike"
  batch_group_name         = "TestBatchGroup"

  release_criteria {
    message_count = 80
  }
}
