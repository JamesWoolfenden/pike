resource "azurerm_chaos_studio_experiment" "pike_gen" {
  location            = "pike"
  name                = "example"
  resource_group_name = "pike"

  identity {
    type = "SystemAssigned"
  }

  selectors {
    name                    = "Selector1"
    chaos_studio_target_ids = ["pike"]
  }

  steps {
    name = "example"
    branch {
      name = "example"
      actions {
        urn           = "pike"
        selector_name = "Selector1"
        parameters = {
          abruptShutdown = "false"
        }
        action_type = "continuous"
        duration    = "PT10M"
      }
    }
  }
}
