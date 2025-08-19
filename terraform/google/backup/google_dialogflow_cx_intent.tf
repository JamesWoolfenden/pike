resource "google_dialogflow_cx_intent" "pike" {
  parent       = google_dialogflow_cx_agent.pike.id
  display_name = "Example"
  priority     = 1
  description  = "Intent example"
  training_phrases {
    parts {
      text = "training"
    }

    parts {
      text = "phrase"
    }

    parts {
      text = "example"
    }

    repeat_count = 1
  }

  parameters {
    id          = "param1"
    entity_type = "projects/-/locations/-/agents/-/entityTypes/sys.date"
  }

  labels = {
    label1 = "value1",
    label2 = "value2"
  }
}
