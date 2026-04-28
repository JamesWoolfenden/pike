resource "azurerm_data_factory_trigger_tumbling_window" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  start_time      = "2022-09-21T00:00:00Z"
  end_time        = "2022-09-21T08:00:00Z"
  frequency       = "Minute"
  interval        = 15
  delay           = "16:00:00"

  annotations = ["example1", "example2", "example3"]
  description = "example description"

  retry {
    count    = 1
    interval = 30
  }

  pipeline {
    name = "pike"
    parameters = {
      Env = "Prod"
    }
  }

  // Self dependency
  trigger_dependency {
    size   = "24:00:00"
    offset = "-24:00:00"
  }

  additional_properties = {
    foo = "value1"
    bar = "value2"
  }
}
