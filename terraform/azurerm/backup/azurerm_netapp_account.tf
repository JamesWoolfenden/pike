resource "azurerm_netapp_account" "pike_gen" {
  name                = "netappaccount"
  location            = "pike"
  resource_group_name = "pike"

  active_directory {
    username            = "aduser"
    smb_server_name     = "SMBSERVER"
    dns_servers         = ["1.2.3.4"]
    domain              = "westcentralus.com"
    organizational_unit = "OU=FirstLevel"
  }

  identity {
    type = "UserAssigned"
    identity_ids = [
      azurerm_user_assigned_identity.example.id
    ]
  }
}
