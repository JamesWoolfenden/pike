resource "azurerm_gallery_application_version" "pike_gen" {
  name                   = "0.0.1"
  gallery_application_id = "pike"
  location               = "pike"

  manage_action {
    install = "[install command]"
    remove  = "[remove command]"
  }

  source {
    media_link = "pike"
  }

  target_region {
    name                   = "pike"
    regional_replica_count = 1
  }
}
