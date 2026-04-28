resource "azurerm_stream_analytics_output_table" "pike_gen" {
  name                      = "output-to-storage-table"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"
  storage_account_name      = "pike"
  storage_account_key       = "pike"
  table                     = "pike"
  partition_key             = "foo"
  row_key                   = "bar"
  batch_size                = 100
}
