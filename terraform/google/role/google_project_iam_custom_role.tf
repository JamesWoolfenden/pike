resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    //google_alloydb_locations
    "alloydb.locations.list",
    //google_alloydb_supported_database_flags
    "alloydb.supportedDatabaseFlags.list",
    //google_beyondcorp_app_connection
    "beyondcorp.appConnections.get",
    //google_beyondcorp_app_connector
    "beyondcorp.appConnectors.get",
    //google_beyondcorp_app_gateway
    "beyondcorp.appGateways.get"
  ]
}
