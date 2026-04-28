resource "azurerm_spring_cloud_gateway_route_config" "pike_gen" {
  name                    = "example"
  spring_cloud_gateway_id = "pike"
  spring_cloud_app_id     = "pike"
  protocol                = "HTTPS"
  route {
    description            = "example description"
    filters                = ["StripPrefix=2", "RateLimit=1,1s"]
    order                  = 1
    predicates             = ["Path=/api5/customer/**"]
    sso_validation_enabled = true
    title                  = "myApp route config"
    token_relay            = true
    uri                    = "https://www.example.com"
    classification_tags    = ["tag1", "tag2"]
  }
}
