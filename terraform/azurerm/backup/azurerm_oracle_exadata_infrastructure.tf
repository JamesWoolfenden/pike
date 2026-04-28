resource "azurerm_oracle_exadata_infrastructure" "pike_gen" {
  name                 = "example-exadata-infra"
  resource_group_name  = "pike"
  location             = "pike"
  zones                = ["1"]
  display_name         = "example-exadata-infra"
  storage_count        = 3
  compute_count        = 2
  shape                = "Exadata.X11M"
  database_server_type = "X11M"
  storage_server_type  = "X11M-HC"
}
