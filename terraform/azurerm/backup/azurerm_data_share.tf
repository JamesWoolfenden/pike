resource "azurerm_data_share" "pike_gen" {
  name        = "example_dss"
  account_id  = "pike"
  kind        = "CopyBased"
  description = "example desc"
  terms       = "example terms"

  snapshot_schedule {
    name       = "example-ss"
    recurrence = "Day"
    start_time = "2020-04-17T04:47:52.9614956Z"
  }
}
