resource "azurerm_data_factory_custom_dataset" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  type            = "Json"

  linked_service {
    name = "pike"
    parameters = {
      key1 = "value1"
    }
  }

  type_properties_json = <<JSON
{
  "location": {
    "container":"${azurerm_storage_container.example.name}",
    "fileName":"foo.txt",
    "folderPath": "foo/bar/",
    "type":"AzureBlobStorageLocation"
  },
  "encodingName":"UTF-8"
}
JSON

  description = "test description"
  annotations = ["test1", "test2", "test3"]
  folder      = "testFolder"

  parameters = {
    foo = "test1"
    Bar = "Test2"
  }

  additional_properties = {
    foo = "test1"
    bar = "test2"
  }

  schema_json = <<JSON
{
  "type": "object",
  "properties": {
    "name": {
      "type": "object",
      "properties": {
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        }
      }
    },
    "age": {
      "type": "integer"
    }
  }
}
JSON
}
