resource "azurerm_web_pubsub_hub" "pike_gen" {
  name          = "tfex_wpsh"
  web_pubsub_id = "pike"
  event_handler {
    url_template       = "https://test.com/api/{hub}/{event}"
    user_event_pattern = "*"
    system_events      = ["connect", "connected"]
  }

  event_handler {
    url_template       = "https://test.com/api/{hub}/{event}"
    user_event_pattern = "event1, event2"
    system_events      = ["connected"]
    auth {
      managed_identity_id = "pike"
    }
  }

  event_listener {
    system_event_name_filter = ["connected"]
    user_event_name_filter   = ["event1", "event2"]
    eventhub_namespace_name  = "pike"
    eventhub_name            = "pike"
  }

  event_listener {
    system_event_name_filter = ["connected"]
    user_event_name_filter   = ["*"]
    eventhub_namespace_name  = "pike"
    eventhub_name            = "pike"
  }

  event_listener {
    system_event_name_filter = ["connected"]
    user_event_name_filter   = ["event1"]
    eventhub_namespace_name  = "pike"
    eventhub_name            = "pike"
  }

  anonymous_connections_enabled = true

  depends_on = [
    azurerm_web_pubsub.example
  ]
}
