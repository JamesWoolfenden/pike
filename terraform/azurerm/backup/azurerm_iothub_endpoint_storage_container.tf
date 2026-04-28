resource "azurerm_iothub_endpoint_storage_container" "pike_gen" {
  resource_group_name = "pike"
  iothub_id           = "pike"
  name                = "acctest"

  container_name = "acctestcont"

  file_name_format           = "{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}"
  batch_frequency_in_seconds = 60
  max_chunk_size_in_bytes    = 10485760
  encoding                   = "JSON"
}
