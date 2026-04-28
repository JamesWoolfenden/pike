resource "azurerm_elastic_san_volume" "pike_gen" {
  name            = "example-esv"
  volume_group_id = "pike"
  size_in_gib     = 1
}
