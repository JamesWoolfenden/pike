data "azurerm_storage_account_blob_container_sas" "pike" {
  connection_string = "DefaultEndpointsProtocol=https;AccountName=pike;AccountKey=AAAAA/123456789012345678901234567890123456789012345678901234567890==;EndpointSuffix=core.windows.net"
  container_name    = "pike"
  https_only        = true

  ip_address = "168.1.5.65"

  start  = "2018-03-21"
  expiry = "2018-03-21"

  permissions {
    read   = true
    add    = true
    create = false
    write  = false
    delete = true
    list   = true
  }

  cache_control       = "max-age=5"
  content_disposition = "inline"
  content_encoding    = "deflate"
  content_language    = "en-US"
  content_type        = "application/json"
}


output "sas" {
  value     = data.azurerm_storage_account_blob_container_sas.pike
  sensitive = true
}
