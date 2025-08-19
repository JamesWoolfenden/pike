resource "google_dialogflow_cx_entity_type" "pike" {
  parent       = google_dialogflow_cx_agent.pike.id
  display_name = "MyEntity"
  kind         = "KIND_MAP"
  entities {
    value    = "value1"
    synonyms = ["synonym1", "synonym2"]
  }
  entities {
    value    = "value2"
    synonyms = ["synonym3", "synonym4"]
  }
  enable_fuzzy_extraction = false
}
