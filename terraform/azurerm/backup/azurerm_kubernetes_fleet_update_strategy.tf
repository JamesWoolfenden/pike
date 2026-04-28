resource "azurerm_kubernetes_fleet_update_strategy" "pike_gen" {
  name                        = "example"
  kubernetes_fleet_manager_id = "pike"
  stage {
    name = "example-stage-1"
    group {
      name = "example-group-1"
    }
    after_stage_wait_in_seconds = 21
  }
}
