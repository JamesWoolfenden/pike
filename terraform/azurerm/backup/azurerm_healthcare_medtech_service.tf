resource "azurerm_healthcare_medtech_service" "pike_gen" {
  name         = "examplemed"
  workspace_id = "pike"
  location     = "east us"

  identity {
    type = "SystemAssigned"
  }

  eventhub_namespace_name      = "example-eventhub-namespace"
  eventhub_name                = "example-eventhub"
  eventhub_consumer_group_name = "$Default"

  device_mapping_json = jsonencode({
    "templateType" : "CollectionContent",
    "template" : [
      {
        "templateType" : "JsonPathContent",
        "template" : {
          "typeName" : "heartrate",
          "typeMatchExpression" : "$..[?(@heartrate)]",
          "deviceIdExpression" : "$.deviceid",
          "timestampExpression" : "$.measurementdatetime",
          "values" : [
            {
              "required" : "true",
              "valueExpression" : "$.heartrate",
              "valueName" : "hr"
            }
          ]
        }
      }
    ]
  })
}
