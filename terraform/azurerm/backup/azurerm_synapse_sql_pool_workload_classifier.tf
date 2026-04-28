resource "azurerm_synapse_sql_pool_workload_classifier" "pike_gen" {
  name              = "example"
  workload_group_id = "pike"

  context     = "example_context"
  end_time    = "14:00"
  importance  = "high"
  label       = "example_label"
  member_name = "dbo"
  start_time  = "12:00"
}
