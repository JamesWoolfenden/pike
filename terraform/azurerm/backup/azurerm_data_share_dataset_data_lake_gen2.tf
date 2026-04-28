resource "azurerm_data_share_dataset_data_lake_gen2" "pike_gen" {
  name               = "accexample-dlg2ds"
  share_id           = "pike"
  storage_account_id = "pike"
  file_system_name   = "pike"
  file_path          = "myfile.txt"
  depends_on = [
    azurerm_role_assignment.example,
  ]
}
