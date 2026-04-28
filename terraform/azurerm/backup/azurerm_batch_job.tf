resource "azurerm_batch_job" "pike_gen" {
  name          = "examplejob"
  batch_pool_id = "pike"
}
