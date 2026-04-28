resource "azurerm_batch_certificate" "pike_gen" {
  resource_group_name  = "pike"
  account_name         = "pike"
  certificate          = filebase64("certificate.pfx")
  format               = "Pfx"
  thumbprint           = "42C107874FD0E4A9583292A2F1098E8FE4B2EDDA"
  thumbprint_algorithm = "SHA1"
}
