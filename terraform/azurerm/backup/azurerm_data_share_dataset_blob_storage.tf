resource "azurerm_data_share_dataset_blob_storage" "pike_gen" {
  name           = "example-dsbsds-file"
  data_share_id  = "pike"
  container_name = "pike"
  storage_account {
    name                = "pike"
    resource_group_name = "pike"
    subscription_id     = "00000000-0000-0000-0000-000000000000"
  }
  file_path = "myfile.txt"
  depends_on = [
    azurerm_role_assignment.example,
  ]
}
