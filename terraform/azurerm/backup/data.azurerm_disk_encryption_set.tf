data "azurerm_disk_encryption_set" "pike_gen" {
  name                = "example-des"
  resource_group_name = "example-resources"
}
