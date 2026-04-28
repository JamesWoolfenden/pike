resource "azurerm_storage_mover_job_definition" "pike_gen" {
  name                     = "example-sjd"
  storage_mover_project_id = "pike"
  agent_name               = "pike"
  copy_mode                = "Additive"
  source_name              = "pike"
  source_sub_path          = "/"
  target_name              = "pike"
  target_sub_path          = "/"
  description              = "Example Job Definition Description"
}
