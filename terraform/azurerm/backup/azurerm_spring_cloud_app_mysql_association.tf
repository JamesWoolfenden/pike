resource "azurerm_spring_cloud_app_mysql_association" "pike_gen" {
  name                = "example-bind"
  spring_cloud_app_id = "pike"
  mysql_server_id     = "pike"
  database_name       = "pike"
  username            = "pike"
}
