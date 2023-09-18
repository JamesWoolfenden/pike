data "azurerm_storage_account_sas" "pike" {
  connection_string = "DefaultEndpointsProtocol=https;AccountName=pike;AccountKey=AAAAA/123456789012345678901234567890123456789012345678901234567890==;EndpointSuffix=core.windows.net"
  https_only        = true
  signed_version    = "2017-07-29"

  resource_types {
    service   = true
    container = false
    object    = false
  }

  services {
    blob  = true
    queue = false
    table = false
    file  = false
  }

  start  = "2018-03-21T00:00:00Z"
  expiry = "2020-03-21T00:00:00Z"

  permissions {
    read    = true
    write   = true
    delete  = false
    list    = false
    add     = true
    create  = true
    update  = false
    process = false
    tag     = false
    filter  = false
  }
}

output "sas2" {
  value     = data.azurerm_storage_account_sas.pike
  sensitive = true
}
