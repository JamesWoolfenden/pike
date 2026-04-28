resource "azurerm_spring_cloud_build_pack_binding" "pike_gen" {
  name                    = "example"
  spring_cloud_builder_id = "pike"
  binding_type            = "ApplicationInsights"
  launch {
    properties = {
      abc           = "def"
      any-string    = "any-string"
      sampling-rate = "12.0"
    }

    secrets = {
      connection-string = "XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX"
    }
  }
}
