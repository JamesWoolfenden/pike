resource "azurerm_kubernetes_fleet_update_run" "pike_gen" {
  name                        = "example"
  kubernetes_fleet_manager_id = "pike"
  managed_cluster_update {
    upgrade {
      type               = "Full"
      kubernetes_version = "1.27"
    }
    node_image_selection {
      type = "Latest"
    }
  }
  stage {
    name = "example"
    group {
      name = "example-group"
    }
    after_stage_wait_in_seconds = 21
  }
}
