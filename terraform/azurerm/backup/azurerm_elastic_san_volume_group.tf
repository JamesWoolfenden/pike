resource "azurerm_elastic_san_volume_group" "pike_gen" {
  name            = "example-esvg"
  elastic_san_id  = "pike"
  encryption_type = "EncryptionAtRestWithCustomerManagedKey"

  encryption {
    key_vault_key_id          = "pike"
    user_assigned_identity_id = "pike"
  }

  identity {
    type         = "UserAssigned"
    identity_ids = ["pike"]
  }

  network_rule {
    subnet_id = "pike"
    action    = "Allow"
  }
}
