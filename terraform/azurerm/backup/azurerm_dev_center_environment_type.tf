resource "azurerm_dev_center_environment_type" "pike_gen" {
  name          = "example-dcet"
  dev_center_id = "pike"

  tags = {
    Env = "Test"
  }
}
