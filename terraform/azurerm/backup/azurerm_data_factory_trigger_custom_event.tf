resource "azurerm_data_factory_trigger_custom_event" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  eventgrid_topic_id  = "pike"
  events              = ["event1", "event2"]
  subject_begins_with = "abc"
  subject_ends_with   = "xyz"

  annotations = ["example1", "example2", "example3"]
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
