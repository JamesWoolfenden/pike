resource "google_sourcerepo_repository" "pike" {
  name = "pike"
  pubsub_configs {
    topic                 = google_pubsub_topic.pike.id
    message_format        = "PROTOBUF"
    service_account_email = "pike-361314@appspot.gserviceaccount.com"
  }
}
