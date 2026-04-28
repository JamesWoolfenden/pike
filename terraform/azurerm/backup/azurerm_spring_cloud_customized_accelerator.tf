resource "azurerm_spring_cloud_customized_accelerator" "pike_gen" {
  name                        = "example"
  spring_cloud_accelerator_id = "pike"

  git_repository {
    url                 = "https://github.com/Azure-Samples/piggymetrics"
    git_tag             = "spring.version.2.0.3"
    interval_in_seconds = 100
  }

  accelerator_tags = ["tag-a", "tag-b"]
  description      = "example description"
  display_name     = "example name"
  icon_url         = "https://images.freecreatives.com/wp-content/uploads/2015/05/smiley-559124_640.jpg"
}
