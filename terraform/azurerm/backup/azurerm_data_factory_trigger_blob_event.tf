resource "azurerm_data_factory_trigger_blob_event" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  storage_account_id  = "pike"
  events              = ["Microsoft.Storage.BlobCreated", "Microsoft.Storage.BlobDeleted"]
  blob_path_ends_with = ".txt"
  ignore_empty_blobs  = true
  activated           = true

  annotations = ["test1", "test2", "test3"]
  description = "example description"

  pipeline {
    name = "pike"
    parameters = {
      Env = "Prod"
    }
  }

  additional_properties = {
    foo = "foo1"
    bar = "bar2"
  }
}
