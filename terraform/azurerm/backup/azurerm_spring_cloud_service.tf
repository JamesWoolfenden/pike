resource "azurerm_spring_cloud_service" "pike_gen" {
  name                = "example-springcloud"
  resource_group_name = "pike"
  location            = "pike"
  sku_name            = "S0"

  config_server_git_setting {
    uri          = "https://github.com/Azure-Samples/piggymetrics"
    label        = "config"
    search_paths = ["dir1", "dir2"]
  }

  trace {
    sample_rate = 10.0
  }

  tags = {
    Env = "staging"
  }
}
