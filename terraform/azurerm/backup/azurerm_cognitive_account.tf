resource "azurerm_cognitive_account" "pike" {
  name                = var.account.name
  location            = "UK South"
  resource_group_name = "pike"
  kind                = var.account.kind


  dynamic_throttling_enabled    = false
  public_network_access_enabled = true
  sku_name                      = var.account.sku_name
  fqdns                         = try(var.account.fqdns)

  identity {
    type = "SystemAssigned"
  }


  local_auth_enabled = true
  tags = {
    name        = "pike"
    permissions = "controller"
  }
}





variable "account" {
  type = object({
    name                       = string
    kind                       = string
    sku_name                   = string
    dynamic_throttling_enabled = bool
    fqdns                      = list(string)
    custom_subdomain_name      = string
  })

  default = {
    name                       = "pike"
    kind                       = "Face"
    sku_name                   = "S0"
    dynamic_throttling_enabled = false
    fqdns                      = []
    custom_subdomain_name      = ""
  }
  description = "(optional) describe your variable"
}
