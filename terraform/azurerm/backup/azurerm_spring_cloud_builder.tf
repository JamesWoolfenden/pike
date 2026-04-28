resource "azurerm_spring_cloud_builder" "pike_gen" {
  name                    = "example"
  spring_cloud_service_id = "pike"

  build_pack_group {
    name           = "mix"
    build_pack_ids = ["tanzu-buildpacks/java-azure"]
  }

  stack {
    id      = "io.buildpacks.stacks.bionic"
    version = "base"
  }
}
