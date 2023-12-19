data "azurerm_batch_certificate" "pike" {
  resource_group_name = "pike"
  name                = "SHA1-42C107874FD0E4A9583292A2F1098E8FE4B2EDDA"
  account_name        = "pike"
}
