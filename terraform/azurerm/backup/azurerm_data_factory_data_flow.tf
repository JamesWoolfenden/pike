resource "azurerm_data_factory_data_flow" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"

  source {
    name = "source1"

    flowlet {
      name = "pike"
      parameters = {
        "Key1" = "value1"
      }
    }

    dataset {
      name = "pike"
    }
  }

  sink {
    name = "sink1"

    flowlet {
      name = "pike"
      parameters = {
        "Key1" = "value1"
      }
    }

    dataset {
      name = "pike"
    }
  }

  script = <<EOT
source(
  allowSchemaDrift: true,
  validateSchema: false,
  limit: 100,
  ignoreNoFilesFound: false,
  documentForm: 'documentPerLine') ~> source1
source1 sink(
  allowSchemaDrift: true,
  validateSchema: false,
  skipDuplicateMapInputs: true,
  skipDuplicateMapOutputs: true) ~> sink1
EOT
}
