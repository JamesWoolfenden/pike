resource "azurerm_mongo_cluster_user" "pike_gen" {
  object_id              = "pike"
  mongo_cluster_id       = "pike"
  identity_provider_type = "MicrosoftEntraID"
  principal_type         = "servicePrincipal"

  role {
    database = "admin"
    name     = "root"
  }
}
