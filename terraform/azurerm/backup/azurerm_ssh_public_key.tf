resource "azurerm_ssh_public_key" "pike" {
  name                = "pike-ssh-key"
  resource_group_name = "pike"
  location            = "uksouth"
  public_key          = tls_private_key.pike.public_key_openssh

  tags = {
    pike = "permissions"
  }
}
