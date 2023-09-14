resource "google_pubsub_topic" "pike" {
  name = "pike"
  labels = {
    pike = "permissions"
  }

  message_retention_duration = "86600s"
}
