resource "azurerm_managed_devops_pool" "pike_gen" {
  name                  = "example-manageddevopspools"
  resource_group_name   = "pike"
  location              = "pike"
  dev_center_project_id = "pike"
  maximum_concurrency   = 1

  azure_devops_organization {
    organization {
      parallelism = 1
      url         = "https://dev.azure.com/example"
    }
  }

  stateless_agent {}

  virtual_machine_scale_set_fabric {
    sku_name = "Standard_D2ads_v5"

    image {
      well_known_image_name = "ubuntu-24.04/buffer"
    }
  }
}
