resource "azurerm_storage_share" "pike_gen" {
  name               = "sharename"
  storage_account_id = "pike"
  quota              = 50

  acl {
    id = "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"

    access_policy {
      permissions = "rwdl"
      start       = "2019-07-02T09:38:21Z"
      expiry      = "2019-07-02T10:38:21Z"
    }
  }
}
